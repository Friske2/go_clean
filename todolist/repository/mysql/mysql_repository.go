package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
	"go_clean/domain"
)

type mysqlTodolistRepo struct {
	DB *sql.DB
}

func NewMysqlAuthorRepository(db *sql.DB) domain.TodolistRepository {
	return &mysqlTodolistRepo{
		DB: db,
	}
}
func (m *mysqlTodolistRepo) Delete(ctx context.Context, id int) error {

	query := `delete from todolist where id = ?`
	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affect != 1 {
		return fmt.Errorf("Weird  Behavior. Total Affected: %d", affect)
	}

	return nil

}

func (m *mysqlTodolistRepo) Update(ctx context.Context, t *domain.TodoList) error {
	query := `update todolist set name=?, description=?,updated_at=now() where id = ?`
	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	res, err := stmt.ExecContext(ctx, t.Name, t.Description, t.ID)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affect != 1 {
		return fmt.Errorf("Weird  Behavior. Total Affected: %d", affect)
	}
	return nil
}

func (m *mysqlTodolistRepo) Insert(ctx context.Context, t *domain.TodoList) error {
	query := `INSERT INTO todolist (name,description,created_at,updated_at) value (?,?,now(),now())`
	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, t.Name, t.Description)
	if err != nil {
		return err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return nil
	}

	t.ID = lastId
	return nil
}

func (m *mysqlTodolistRepo) GetOne(ctx context.Context, query string, args ...interface{}) (res domain.TodoList, err error) {
	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		return domain.TodoList{}, err
	}

	row := stmt.QueryRowContext(ctx, args...)
	res = domain.TodoList{}

	err = row.Scan(&res.ID, &res.Name, &res.Description, &res.UpdatedAt, &res.CreatedAt)
	return res, err
}

func (m *mysqlTodolistRepo) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.TodoList, err error) {
	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()

	result = make([]domain.TodoList, 0)
	for rows.Next() {
		t := domain.TodoList{}
		err = rows.Scan(
			&t.ID,
			&t.Name,
			&t.Description,
			&t.UpdatedAt,
			&t.CreatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}
	return result, nil
}
func (m *mysqlTodolistRepo) Get(ctx context.Context) (res []domain.TodoList, err error) {
	query := `select * from todolist`
	res, err = m.fetch(ctx, query)
	return
}
func (m *mysqlTodolistRepo) GetByID(ctx context.Context, id int) (domain.TodoList, error) {
	query := `select * from todolist where id=?`
	res, err := m.GetOne(ctx, query, id)
	if err != nil {
		return domain.TodoList{}, err
	}
	return res, nil
}
