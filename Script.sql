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
values 	('Mirza', 'Mengambil Botol Minum', '2024-04-29 15:05:36.208328'),
		('Epul', 'Mengambil Hati Wanita Tanpa Bilang Bilang', '2024-03-15 16:05:36.208328' ),
		('Herdi', 'Melakukan pencucian uang', '2024-06-20 10:05:36.208328'),
		('Atthar', 'Melakukan kecurangan saat bermain FIFA', '2024-07-25 17:05:36.208328'),
		('Ahmad', 'Melakukan Pelanggaran dengan memburu hewan yang di lindungi saat mendaki gunung (Perpu Kementrian Kehutanan)', '2024-04-19 15:05:36.208328')
		
--case1		
select * from criminal_records where crime_status = 'Open' and crime_date > '2024-03-15' 

--case2
select * from criminal_records order by crime_date desc limit 3

--case3
select count(*) from criminal_records where crime_status in ('Open')

--case4
select count(*) from criminal_records where crime_status in ('Close')

--case5
select * from criminal_records  where crime_description ilike '%mengambil%'		

---When the suspect is found to be innocent, update the criminal data to 'Close'

select id, suspect_name from criminal_records

UPDATE criminal_records
SET crime_status =
    CASE 
        WHEN suspect_name = 'Herdi' OR suspect_name = 'Ahmad' THEN 'Open'
        ELSE 'Close'
    END;




---When the suspect is truly innocent and his status has become 'Close', then remove him from the criminal list

delete from criminal_records 
where crime_status = 'Close';






---reset---
SELECT pg_get_serial_sequence('criminal_records', 'id');

ALTER SEQUENCE public.criminal_records_id_seq RESTART WITH 1;

delete from criminal_records 


insert into
	criminal_records (suspect_name,
	crime_description,
	crime_date )
values 	('Mirza', 'Mengambil Botol Minum', '2024-04-29 15:05:36.208328'),
		('Epul', 'Mengambil Hati Wanita Tanpa Bilang Bilang', '2024-03-15 16:05:36.208328' ),
		('Herdi', 'Melakukan pencucian uang', '2024-06-20 10:05:36.208328'),
		('Atthar', 'Melakukan kecurangan saat bermain FIFA', '2024-07-25 17:05:36.208328'),
		('Ahmad', 'Melakukan Pelanggaran dengan memburu hewan yang di lindungi saat mendaki gunung (Perpu Kementrian Kehutanan)', '2024-04-19 15:05:36.208328')

select * from criminal_records
