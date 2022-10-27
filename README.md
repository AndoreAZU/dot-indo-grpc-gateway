
# Sample Project GRPC - Gateway

Project ini dibangun menggunakan freamwork GRPC-Gateway. Adapun alasan mengapa menggunakan GRPC-Gateway adalah:
* Tidak perlu lagi memusingkan bagaimana handling panic / error secara global. Karna GRPC-Gateway ini sudah secara otomatis melakukan panic handling secara global.
* Mempermudah pembuatan project api yang membutuhkan waktu cepat. Bagi saya dengan membuat file protobuf akan sangat memudahkan development backend, karna secara tidak langsung kita sudah mendesign bagaimana struktur api nantinya, sehingga kita tinggal perlu memikirkan core logic nya saja.
* Dengan adanya file protobuf, kita bisa secara otomatis membuat file JSON untuk dokumentasi OpenAPI sekaligus.
* Sangat berguna apabila kita membutuhkan 2 protocol api sekaligus, GRPC dan Rest HTTP

Dan mungkin masih ada beberapa keuntungan lainnya yang mungkin dirasakan pengguna lainnya. Tapi selain itu, adapun beberapa kelemahan dari GRPC-Gateway ini dalam pandangan saya.
* Membutuhkan cost resource untuk transform grpc ke rest. Karna sejatinya GRPC-Gateway ini bertindak seperti reverse proxy, sehingga akan ada cost untuk melakukan transform request di setiap transaction api.
* Cukup sulit untuk mendesign project repo diawal, karna memang belum ada design yang resmi digunakan. Namun jika sudah terbiasa dan memiliki design tersendiri, itu akan memudahkan kita.
* Karna untuk mendesign request response api menggunakan protobuf, maka kita tidak bisa membuat 1 generic / general response. Untuk beberapa developer mungkin hal ini akan menyulitkan, karna akan lebih mudah jika kita bisa membuat 1 general model request response.

Dan juga masih ada kekurangan lain yang mungkin developer lain rasakan.



## Run Lokal

Clone the project

```bash
  git clone https://github.com/AndoreAZU/dot-indo-grpc-gateway.git
```

Go to the project directory

```bash
  cd ${project-dir}
```

Install dependencies

```bash
  make install
```

Start the server

```bash
  make run env={profile}
```
## Tech Stack

**Backend:** Golang, GRPC-Gateway, Postgress, Redis

## Deployment

Untuk mendeploy ke server, kita bisa menggunakan 2 cara. Dengan build file bin dan menggunakan container.

First time, we need install all depedency
```bash
  make install
```

Then, if we want to run with container
```bash
  profile=${profile} make deploy-container
```

Without container
```bash
  profile=${profile} make deploy
```
## Badges

[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)
[![GPLv3 License](https://img.shields.io/badge/License-GPL%20v3-yellow.svg)](https://opensource.org/licenses/)
[![AGPL License](https://img.shields.io/badge/license-AGPL-blue.svg)](http://www.gnu.org/licenses/agpl-3.0)

