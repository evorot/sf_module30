package storage

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Task struct {
	Id         int
	Opened     int64
	Closed     int64
	AuthorID   int
	AssignedID int
	Title      string
	Content    string
}

type Storage struct {
	db *pgxpool.Pool
}

func New(dbURL string) (*Storage, error) {
	db, err := pgxpool.Connect(context.Background(), dbURL)
	if err != nil {
		return nil, err
	}
	return &Storage{db: db}, nil
}

func (s *Storage) NewTask(t Task) (int, error) {
	var id int
	err := s.db.QueryRow(context.Background(),
		"INSERT INTO tasks (author_id, assigned_id, title, content) "+
			"VALUES ($1,$2,$3,$4) RETURNING tasks.id;",
		t.AuthorID, t.AssignedID, t.Title, t.Content).Scan(&id)
	return id, err
}

func (s *Storage) AllTasks() ([]Task, error) {
	rows, err := s.db.Query(context.Background(),
		"SELECT * FROM tasks;")
	if err != nil {
		return nil, err
	}
	var tasks []Task
	for rows.Next() {
		var t Task
		err = rows.Scan(&t.Id, &t.Opened, &t.Closed, &t.AuthorID, &t.AssignedID, &t.Title, &t.Content)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, rows.Err()
}

func (s *Storage) Task(taskID int) (Task, error) {
	var t Task
	err := s.db.QueryRow(context.Background(),
		"SELECT * FROM tasks WHERE id=$1;", taskID).Scan(&t.Id, &t.Opened, &t.Closed, &t.AuthorID, &t.AssignedID, &t.Title, &t.Content)
	return t, err
}

func (s *Storage) AuthorsTasks(authorID int) ([]Task, error) {
	rows, err := s.db.Query(context.Background(),
		"SELECT * FROM tasks WHERE tasks.author_id=$1;", authorID)
	if err != nil {
		return nil, err
	}
	var tasks []Task
	for rows.Next() {
		var t Task
		err = rows.Scan(&t.Id, &t.Opened, &t.Closed, &t.AuthorID, &t.AssignedID, &t.Title, &t.Content)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, rows.Err()
}

func (s *Storage) LabelsTasks(labelID int) ([]Task, error) {
	rows, err := s.db.Query(context.Background(),
		"SELECT tasks.id, tasks.opened, tasks.closed, tasks.author_id, tasks.assigned_id, title, content "+
			"FROM tasks,tasks_labels "+
			"WHERE label_id = $1 AND tasks.id=tasks_labels.task_id;", labelID)
	if err != nil {
		return nil, err
	}
	var tasks []Task
	for rows.Next() {
		var t Task
		err = rows.Scan(&t.Id, &t.Opened, &t.Closed, &t.AuthorID, &t.AssignedID, &t.Title, &t.Content)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, rows.Err()
}

func (s *Storage) DeleteTask(taskID int) (bool, error) {
	ct, err := s.db.Exec(context.Background(),
		"DELETE FROM tasks WHERE id=$1;", taskID)
	if err != nil {
		return ct.Delete(), err
	}
	return ct.Delete(), nil
}

func (s *Storage) UpdateTask(taskID int, t Task) (bool, error) {
	ct, err := s.db.Exec(context.Background(),
		"UPDATE tasks "+
			"SET author_id = $1, "+
			"assigned_id = $2, "+
			"title = $3, "+
			"content = $4 "+
			"WHERE id = $5", t.AuthorID, t.AssignedID, t.Title, t.Content, taskID)
	if err != nil {
		return ct.Update(), err
	}
	return ct.Update(), nil
}
