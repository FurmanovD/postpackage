-- +goose Up
-- +goose StatementBegin
CREATE DATABASE IF NOT EXISTS `default` /*!40100 DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci */;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `goose_db_version`;
-- +goose StatementEnd

