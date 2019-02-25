# Go-Simple-Api
Simple Api With Golang


==================
Query Create Table 
==================

CREATE TABLE `mahasiswa` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(45) DEFAULT NULL,
  `phone` varchar(45) DEFAULT NULL,
  `address` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=latin1;


- Endpoint :
    - http://localhost:9090/insertmhs

        raw body :
        {
            "name" : "ethant hunt",
            "phone" : "987654321",
            "address": "New york America"
        }

    - http://localhost:9090/getmhs

    - http://localhost:9090/updatemhs

        raw body :
        {
            "id" : 2,
            "name" : "fifi",
            "phone" : "123456789",
            "address": "sukun pondok indah bandung"
        }