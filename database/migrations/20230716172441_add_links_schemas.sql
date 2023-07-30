-- +goose Up
-- +goose StatementBegin
CREATE TABLE link_type (
  id INT NOT NULL AUTO_INCREMENT,
  name VARCHAR(50) NOT NULL,
  icon_class VARCHAR(250) NOT NULL,
  PRIMARY KEY(id)
);

-- +goose StatementEnd
-- +goose StatementBegin
CREATE TABLE link (
  id INT NOT NULL AUTO_INCREMENT,
  user_id INT NOT NULL,
  url VARCHAR(255) NOT NULL,
  link_type_id INT NOT NULL,
  published BOOLEAN NOT NULL,
  PRIMARY KEY(id),
  CONSTRAINT fk_link_type FOREIGN KEY (link_type_id) REFERENCES link_type(id),
  CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES user(id)
);

-- +goose StatementEnd

-- +goose Down

-- +goose StatementBegin
DROP TABLE IF EXISTS link;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS link_type;
-- +goose StatementEnd