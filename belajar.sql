select * from pg_tables where schemaname='belajar'

---Create table---
CREATE TABLE barang 
(
    kode INT NOT NULL,
    name VARCHAR(100) NOT NULL,
    harga INT NOT NULL DEFAULT 1000,
    jumlah INT NOT NULL DEFAULT 0,
    "waktu dibuat" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

select * from barang; 


---Create table---
create table products
(
	id VARCHAR(10) not null,
	name VARCHAR(100) not null,
	description text,
	price INT not null,
	quantity INT not null default 0, 
	created_at TIMESTAMP not null default CURRENT_TIMESTAMP
)

---Insert---
INSERT into products(id, name, price, quantity)
values('P0001', 'Mie Ayam Original', 15000, 100);

INSERT into products(id, name, description, price, quantity)
values('P0002', 'Mie Ayam Bakso Tahu', 'Mie Ayam Original + Bakso Tahu', 20000, 100);

INSERT into products (id, name, price, quantity)
values ('P0003', 'Mie Ayam Ceker', 20000, 100),
		('P0004', 'Mie Ayam Special', 25000, 100),
		('P0005', 'Mie Ayam Yamin', 15000, 100);
	
select * from products;

select id, name, price, quantity from products;

alter table products
add primary key (id);

select
	id,
	name,
	price,
	quantity
from
	products
where
	price = 20000;


CREATE TYPE PRODUCT_CATEGORY AS ENUM ('Makanan', 'Minuman', 'Lain-lain')

ALTER TABLE products
ADD COLUMN category PRODUCT_CATEGORY;

select * from products;


---Update---

update products 
set category = 'Makanan'
where id = 'P0001'

update products 
set category = 'Makanan'
where id = 'P0002'

update products 
set category = 'Makanan'
where id = 'P0003'

update products 
set category = 'Makanan'
where id = 'P0004'

update products 
set category = 'Makanan'
where id = 'P0005'


update products 
set description = 'Mie Ayam + Ceker',
	price = price + 15000
where id = 'P0003'

update products 
set price = price + 5000
where id = 'P0004'


select * from products;


insert into products (id, name, price, quantity, category)
values ('P0009', 'Contoh', 10000, 100, 'Minuman')



---delete---
delete from products 
where id = 'P0009'





-----operator perbandingan-----


select * from products;

select * from products where price != 20000

select * from products where price > 15000

select * from products where price <= 15000


insert into products(id, name, price, quantity, category)
values 	('P0006', 'Es teh tawar', 10000, 100, 'Minuman'),
		('P0007', 'Es Campur', 20000, 100, 'Minuman'),
		('P0008', 'Es Jeruk', 10000, 100, 'Minuman');

-----And & Or Operator-----

---And---

select * from products 
	
select * from products
where price > 15000
and category = 'Makanan'

---Or---
select * from products
where price < 20000
or category = 'Minuman'


---Prioritas ()---
select
	*
from
	products
where
	(price < 10000)
	and quantity > 100
	or category = 'Makanan'
	
---Like Operator---
select
	*
from
	products
where
	name ilike '%mie%'

	
	
---Null Operator---
select
	*
from
	products
where
	description isnull 

---notnull operator---	
select
	*
from
	products
where
	description notnull
	
	
---between operator---
select * 
from products
where price between 15000 and 20000

select *
from products r  
where r.description between '' and ''

---In operator---
select * 
from products
where category in ('Makanan')



---mengurutkan data---

select * 
from products
order by
	description asc;

select * from products;


---limit clause---
select * 
from products p 
where price  > 0 
order by price asc, id desc;

select * 
from products p 
where price  > 0 
order by price asc, id desc limit 2;

select * 
from products p 
where price  > 0 
order by price asc, id desc limit 2 offset 2;


---Flow Control Case---
select id, name, 
	case category 
		when 'Makanan' then 'Enak'
		when 'Minuman' then 'Segar'
		else 'Apa itu'
		end as category
from products 

select * from products