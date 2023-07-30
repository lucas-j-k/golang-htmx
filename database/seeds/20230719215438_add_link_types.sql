-- +goose Up
-- +goose StatementBegin
INSERT INTO
  link_type (name, icon_class)
VALUES
  ("Facebook", "fa-brands fa-square-facebook"),
  ("Github", "fa-brands fa-github"),
  ("LinkedIn", "fa-brands fa-linkedin"),
  ("Youtube", "fa-brands fa-youtube"),
  ("Other", "fa-solid fa-link");

-- +goose StatementEnd