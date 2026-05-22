package vacancy

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type VacancyRepository struct {
	DbPool *pgxpool.Pool
	Logger *zerolog.Logger
}

func NewVacancyRepository(dbpool *pgxpool.Pool, logger *zerolog.Logger) *VacancyRepository {
	repo := VacancyRepository{
		DbPool: dbpool,
		Logger: logger,
	}
	return &repo
}

func (r *VacancyRepository) CountAll() int {
	query := "SELECT count(*) FROM vacancies"
	var count int
	r.DbPool.QueryRow(context.Background(), query).Scan(&count)
	return count
}

func (r *VacancyRepository) GetAll(limit, offset int) ([]Vacancy, error) {
	query := "SELECT * FROM vacancies ORDER BY createdat LIMIT @limit OFFSET @offset"
	args := pgx.NamedArgs{
		"limit":  limit,
		"offset": offset,
	}
	rows, err := r.DbPool.Query(context.Background(), query, args)
	if err != nil {
		return nil, err
	}
	vacancies, err := pgx.CollectRows(rows, pgx.RowToStructByName[Vacancy])

	if err != nil {
		return nil, err
	}
	return vacancies, nil
}

func (r *VacancyRepository) AddVacancy(form VacancyCreateForm) error {
	query := "INSERT INTO vacancies ( role ,companyType,location,companyName ,salary , email ) VALUES ( @role ,@companyType,@location,@companyName ,@salary , @email)"
	args := pgx.NamedArgs{
		"role":        form.Role,
		"companyType": form.CompanyType,
		"location":    form.Location,
		"companyName": form.CompanyName,
		"salary":      form.Salary,
		"email":       form.Email,
	}
	_, err := r.DbPool.Exec(context.Background(), query, args)
	if err != nil {
		return fmt.Errorf("Невозможно создать выкансию:%w", err)
	}
	return nil
}
