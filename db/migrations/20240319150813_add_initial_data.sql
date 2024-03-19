-- +goose Up
-- +goose StatementBegin
INSERT INTO `packages` (`id`, `name`, `items_per_package`) VALUES
(1, 'XS', 250),
(2, 'S', 500),
(3, 'M', 1000),
(4, 'L', 2000),
(5, 'XL', 5000);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM `packages` WHERE `id` IN (1, 2, 3, 4, 5);
-- +goose StatementEnd

