-- Users
-- password: "password123" (bcrypt hash)
INSERT INTO users (id, name, email, phone, password, role) VALUES
(1, 'Admin', 'admin@rentbike.com', '081234567893', '$2a$10$DkBRkLG7u5VDBvWK2RQFMe4xcxzBbQ7cLNjMmNwa14HFcR5XqrlMW', 'admin'),
(2, 'Budi Santoso', 'budi@rentbike.com', '081234567890', '$2a$10$DkBRkLG7u5VDBvWK2RQFMe4xcxzBbQ7cLNjMmNwa14HFcR5XqrlMW', 'customer'),
(3, 'Siti Nurhaliza', 'siti@rentbike.com', '081234567891', '$2a$10$DkBRkLG7u5VDBvWK2RQFMe4xcxzBbQ7cLNjMmNwa14HFcR5XqrlMW', 'customer'),
(4, 'Ahmad Rizky', 'ahmad@rentbike.com', '081234567892', '$2a$10$DkBRkLG7u5VDBvWK2RQFMe4xcxzBbQ7cLNjMmNwa14HFcR5XqrlMW', 'customer');

-- Vehicles
INSERT INTO vehicles (id, type, brand, name, category, description, price_per_hour, status) VALUES
(1, 'motor', 'Honda', 'Vario 150', 'Matic', 'Motor matic irit dan nyaman untuk harian. Mesin 150cc halus, bagasi luas muat helm.', 50000, 'available'),
(2, 'motor', 'Yamaha', 'NMAX', 'Matic', 'Motor matic premium dengan desain sporty. Mesin 155cc bertenaga, dilengkapi ABS dan lampu LED.', 75000, 'available'),
(3, 'motor', 'Honda', 'Beat Street', 'Matic', 'Motor matic ringan dan lincah. Mesin 110cc sangat irit bahan bakar, cocok untuk mahasiswa.', 35000, 'available'),
(4, 'motor', 'Yamaha', 'XSR 155', 'Sport', 'Motor sport retro bergaya klasik. Mesin 155cc dengan tampilan naked bike yang stylish.', 80000, 'available'),
(5, 'motor', 'Honda', 'CRF 150L', 'Trail', 'Motor trail tangguh untuk petualangan off-road. Suspensi panjang siap menjelajah medan berat.', 65000, 'available'),
(6, 'motor', 'Suzuki', 'Satria F150', 'Sport', 'Motor sport dengan mesin 150cc DOHC. Akselerasi responsif dan handling tajam.', 45000, 'available'),
(7, 'motor', 'Kawasaki', 'Ninja 250', 'Sport', 'Motor sport fairing 250cc dengan tampilan agresif. Performa maksimal untuk touring jarak jauh.', 100000, 'available'),
(8, 'motor', 'Yamaha', 'Aerox 155', 'Matic', 'Motor matic sporty dengan desain agresif. Mesin 155cc VVA bertenaga, cocok untuk anak muda.', 65000, 'available'),
(9, 'mobil', 'Toyota', 'Avanza', 'MPV', 'Mobil keluarga serbaguna dengan kabin luas. Irit bahan bakar dan perawatan mudah.', 150000, 'available'),
(10, 'mobil', 'Honda', 'Brio', 'Hatchback', 'Mobil hatchback kompak dan lincah. Cocok untuk di perkotaan, mudah parkir dan irit bensin.', 120000, 'available'),
(11, 'mobil', 'Daihatsu', 'Xenia', 'MPV', 'Mobil keluarga nyaman dengan fitur lengkap. AC double blower, lega untuk 7 penumpang.', 150000, 'available'),
(12, 'mobil', 'Suzuki', 'Ertiga', 'MPV', 'Mobil MPV stylish dengan performa handal. Kabin senyap dan suspensi empuk untuk perjalanan jauh.', 160000, 'available'),
(13, 'mobil', 'Toyota', 'Fortuner', 'SUV', 'Mobil SUV gagah dan tangguh. Mesin diesel bertenaga, cocok untuk segala medan.', 300000, 'available'),
(14, 'mobil', 'Honda', 'HR-V', 'SUV', 'Mobil SUV kompak dengan desain sporty. Fitur lengkap dan handling responsif.', 250000, 'available'),
(15, 'mobil', 'Mitsubishi', 'Pajero Sport', 'SUV', 'Mobil SUV premium dengan performa off-road. Kabin mewah dan fitur keselamatan lengkap.', 350000, 'available'),
(16, 'mobil', 'Toyota', 'Yaris', 'Hatchback', 'Mobil hatchback modern dengan desain aerodinamis. Handling lincah dan fitur hiburan lengkap.', 130000, 'available');
