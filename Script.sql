select * from products 

CREATE TABLE "criminal_records" (
  "id" serial PRIMARY KEY,
  "suspect_name" varchar NOT NULL,
  "crime_description" text,
  "crime_date" timestamp,
  "crime_status" varchar DEFAULT 'Open'
);

select * from criminal_records 

insert into
	criminal_records (suspect_name,
	crime_description,
	crime_date )
values ('Epul', 'Megambil Hati Wanita Tanpa Bilang Bilang', '2024-03-15 16:05:36.208328' ),
		('Herdi', 'Melakukan pencucian uang', '2024-06-20 10:05:36.208328'),
		('Atthar', 'Melakukan kecurangan saat bermain FIFA', '2024-07-25 17:05:36.208328'),
		('Ahmad', 'Melakukan Pelanggaran dengan memburu hewan yang di lindungi saat mendaki gunung (Perpu Kementrian Kehutanan)', '2024-04-19 15:05:36.208328')


select * from criminal_records where crime_status = 'Open' and crime_date > '2024-03-15' 

select * from criminal_records order by crime_date desc limit 3

select count(*) from criminal_records where crime_status in ('Open')

select count(*) from criminal_records where crime_status in ('Close')

select * from criminal_records  where crime_description ilike '%mengambil%'

