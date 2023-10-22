-- +goose Up
-- +goose StatementBegin
INSERT INTO "product" (id,name, main_warehouse_id, number_of_products_on_main_warehouse)
VALUES (1,'Ноутбук',(SELECT id FROM warehouse where warehouse.name='А'),5);
INSERT INTO "product" (id,name, main_warehouse_id, number_of_products_on_main_warehouse)
VALUES (2,'Телевизор',(SELECT id FROM warehouse where warehouse.name='А'),3);
INSERT INTO "product" (id,name, main_warehouse_id, number_of_products_on_main_warehouse,
                       id_warehouses, number_products_on_warehouse)
VALUES (3,'Телефон',(SELECT id FROM warehouse where warehouse.name='Б'),3,
        ARRAY[(SELECT id FROM warehouse where warehouse.name='З'),
        (SELECT id FROM warehouse where warehouse.name='В')],ARRAY[0,0]);
INSERT INTO "product" (id,name, main_warehouse_id, number_of_products_on_main_warehouse)
VALUES (4,'Системный блок',(SELECT id FROM warehouse where warehouse.name='Ж'),4);
INSERT INTO "product" (id,name, main_warehouse_id, number_of_products_on_main_warehouse,
                       id_warehouses, number_products_on_warehouse)
VALUES (5,'Часы',(SELECT id FROM warehouse where warehouse.name='Ж'),1,
        ARRAY[(SELECT id FROM warehouse where warehouse.name='А')],
        ARRAY[0]);
INSERT INTO "product" (id,name, main_warehouse_id, number_of_products_on_main_warehouse)
VALUES (6,'Микрофон',(SELECT id FROM warehouse where warehouse.name='Ж'),1);



INSERT INTO orders (id,kind_of_tech, number_of_tech) VALUES (10,ARRAY[
     (SELECT id FROM product WHERE name='Ноутбук'),
     (SELECT id FROM product WHERE name='Телефон'),
     (SELECT id FROM product WHERE name='Микрофон')],ARRAY[2,1,1]);
INSERT INTO orders (id, kind_of_tech, number_of_tech) VALUES (
     11,ARRAY[(SELECT id FROM product WHERE name='Телевизор')],
     ARRAY[3]);
INSERT INTO orders (id, kind_of_tech, number_of_tech) VALUES (
     14,ARRAY[(SELECT id FROM product WHERE name='Ноутбук'),
     (SELECT id FROM product WHERE name='Системный блок')],ARRAY[3,4]);
INSERT INTO orders (id, kind_of_tech, number_of_tech) VALUES (
     15,ARRAY[(SELECT id FROM product WHERE name='Часы')],ARRAY[1]);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
