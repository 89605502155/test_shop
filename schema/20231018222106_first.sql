-- +goose Up
CREATE table warehouse (
    id serial PRIMARY KEY,
    name varchar(255) NOT NULL
);
CREATE table product(
    id serial PRIMARY KEY,
    name varchar(255) NOT NULL,
    main_warehouse_id int references warehouse(id) on delete cascade,
    number_of_products_on_main_warehouse int,
    id_warehouses int[],
--         FOREIGN KEY (EACH ELEMENT OF id_warehouses) REFERENCES warehouse,
    number_products_on_warehouse int[]
);
CREATE table orders(
    id serial PRIMARY KEY,
    kind_of_tech int[],
--         FOREIGN KEY (EACH ELEMENT OF kind_of_tech) REFERENCES product,
    number_of_tech int[]
);
-- +goose StatementBegin
-- +goose StatementEnd

-- +goose Down
DROP table orders;
DROP table product;
DROP table warehouse;
-- +goose StatementBegin
-- +goose StatementEnd
