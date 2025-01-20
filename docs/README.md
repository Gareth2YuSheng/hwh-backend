# Homework-Help Backend

CVWO Assignment Project
Server written in GOLANG, Database using POSTGRESQL, Media storage using CLOUDINARY

By Gareth Too Yu Sheng

## Setting Up the Project

### 0. Downloading the Github Repository

If you have already downloaded and unzipped the Github Repo, please skip ahead to the next steps.

To download the Github Repo, at the top of the page select _Code > Download ZIP_ to download the repo.
![download repo](https://github.com/Gareth2YuSheng/hwh-backend/blob/main/docs/readmeImages/download_repo.png)

Save the zip somewhere on your local machine. After downloading, extract the project.
After extracting the project, you should see a folder containing all the source code for the project.
![download repo](https://github.com/Gareth2YuSheng/hwh-backend/blob/main/docs/readmeImages/repo_unzipped.png)

Moving forward I will refer to this folder as the _main project directory_.

### 1. Installing Go

If you already have Golang installed on your machine, feel free to skip this step.

To download Go, go to [the official Go website](https://go.dev/dl/) to download Go.
Select the suitable one for your machine to download, after downloading the installer file, install it.
![go download page](https://github.com/Gareth2YuSheng/hwh-backend/blob/main/docs/readmeImages/go_download.png)

For windows I recommend using the .msi intaller. Inside the installer, leave everything as the default and install.

For a clearer installation and setup process, you can watch [this video](https://www.youtube.com/watch?v=DFiXJKIF2ss).

### 2. Database Setup

If you already have a PostgreSQL database instance setup, feel free to skip this section and use your own database url.

To setup a Local Instance of PostgreSQL, download PostgreSQL and pgAdmin [here](https://www.enterprisedb.com/downloads/postgres-postgresql-downloads).
![pgsql download page](https://github.com/Gareth2YuSheng/hwh-backend/blob/main/docs/readmeImages/pgsql_download.png)

Alternatively if you already have PostgreSQL installed but not pgAdmin, you can download pgAdmin [here](https://www.pgadmin.org/download/).
Select the specific one for your machine and install, when running the PostgreSQL installer, make sure to leave these options checked to
simultaneously install pgAdmin as well.
![pgsql install](https://github.com/Gareth2YuSheng/hwh-backend/blob/main/docs/readmeImages/pgsql_install.png)

Make sure to set up a _root password_ when prompted and you can leave the _port number_ as the default.
![pgsql install2](https://github.com/Gareth2YuSheng/hwh-backend/blob/main/docs/readmeImages/pgsql_install_2.png)
![pgsql install3](https://github.com/Gareth2YuSheng/hwh-backend/blob/main/docs/readmeImages/pgsql_install_3.png)

Go ahead and leave the rest of the options as default as well. If prompted, Stackbuilder is **NOT** required for this project, but you can install it if you'd like.
![pgsql install3](https://github.com/Gareth2YuSheng/hwh-backend/blob/main/docs/readmeImages/pgsql_install_4.png)

Once installed, start up pgAdmin. Click on _Servers_ on the left hand side menu, you will be prompted to enter the _root password_ that you created during the installation.
![pgadmin setup](https://github.com/Gareth2YuSheng/hwh-backend/blob/main/docs/readmeImages/pgadmin_setup.png)

Under _Servers > PostgreSQL > Databases_, right click on _Databases_, click on _Create > Database..._ to create an new database.

![pgadmin setup2](https://github.com/Gareth2YuSheng/hwh-backend/blob/main/docs/readmeImages/pgadmin_setup_2.png)

Name the database whatever you like, no SQL file is required as the server will handle all the table setup once connected to the database.
Setting a password for the database is optional.
![pgadmin setup3](https://github.com/Gareth2YuSheng/hwh-backend/blob/main/docs/readmeImages/pgadmin_setup_3.png)

For a clearer installation and setup process of PostgreSQL and pgAdmin, you can watch [this video](https://www.youtube.com/watch?v=4qH-7w5LZsA).

### 3. Creating the .env file

After setting up the database, create a .env file in the main project directory. It should look like this in the directory.
![env file](https://github.com/Gareth2YuSheng/hwh-backend/blob/main/docs/readmeImages/env_file.png)

Next open the .env file in a text editor, and paste the following template into the file.
![env file2](https://github.com/Gareth2YuSheng/hwh-backend/blob/main/docs/readmeImages/env_file_template.png)

The .env file template can be found at the bottom of this document. Alternatively a template text file for the env file has been provided in the repo.

Now that the local database instance has been created, we can update the `DB_URL` variable in the .env file.
The db url looks something like:
`DB_URL=postgres://[db user]:[db password (optional)]@[db link/endpoint]:[portname]/[db name]?sslmode=disable`
Replace the fields in [] with the information of your database instance. If you didn't change the _db user_ during installation it should be _postgres_.
Replace the _db password_ with whatever password you have set for your database, if you didn't set one you can remove the _:_ from the url and ignore the _db password_ option.
For a local instance the _db endpoint_ should be localhost, the _portname_ if left as the default during installation, it should be _5432_.
If you changed it during installtion, use whatever port that you set. The _db name_ should be whatever you named the database instance in the previous section.
The final `DB_URL` should look something like `postgres://postgres:root@localhost:5432/hwhelp?sslmode=disable`. For local instances please include the `sslmode=disable`.
If your database is hosted online, you can omit this parameter.

### 4. Cloudinary Setup

A Cloudinary account will be required for media storage, not to worry as there is a free tier available.
Sign up for a Cloudinary account [here](https://cloudinary.com/users/register_free).

After signing up for an account, we need to note down your Cloudinary Url for the .env file.
First, login using your new account and go to the bottom left there will be a settings icon above your profile. Click on it to go to the settings page.
At the top of the left side menu, you will see your `CLOUD_NAME`, under _Product environment settings_ navigate to _Upload Presets_.

![upload preset](https://github.com/Gareth2YuSheng/hwh-backend/blob/main/docs/readmeImages/upload_presets.png)

There we will need to add a new _Upload Preset_, click on _Add Upload Preset_ on the top right and name it whatever you like, e.g. hw-help-preset,
give a name for your _Asset Folder_ e.g. hw-help-images and make sure to change the _Signing mode_ to **Unsigned**.
Change the _Generated public ID_ to "Use the filename of the uploaded file as the public ID" and leave the rest of the settings as the default.
![new upload preset](https://github.com/Gareth2YuSheng/hwh-backend/blob/main/docs/readmeImages/new_upload_preset.png)

Once done click Save on the top right. If succesful, you will see a new _upload preset_ added to the list of upload presets.
![upload preset2](https://github.com/Gareth2YuSheng/hwh-backend/blob/main/docs/readmeImages/upload_presets_2.png)

Once the _upload preset_ has been created successfully, update the `CLOUDINARY_UPLOAD_PRESET` in the .env file to whatever you named your _upload preset_.
e.g. For mine it will be `CLOUDINARY_UPLOAD_PRESET=HomeworkHelp`

Next we need the API keys, similarly under _Product environment settings_, navigate to _API Keys_.
![api keys](https://github.com/Gareth2YuSheng/hwh-backend/blob/main/docs/readmeImages/api_keys.png)

Your _Cloudinary Url_ looks like `cloudinary://API_KEY:API_SECRET@CLOUD_NAME`. There will already be 1 API Key created for you by default, it is
up to you whether you wish to use it or create a new one. Select an `API Key` of your choice and copy the `API Key` and `API Secret`
into your _Cloudinary Url_ in the .env file. To copy the `API Secret` you need to click on the eye icon to view it first.
On this page at the top also provides you with a sample Cloudinary Url with your `CLOUD_NAME`.
![api keys2](https://github.com/Gareth2YuSheng/hwh-backend/blob/main/docs/readmeImages/api_key_2.png)

Alternatively you can also copy this into the .env file, but you are will required to copy the `API Key` and `API Secret` manually.
If you are having trouble finding your `API Key`, `API Secret` and `Cloud Name`, you can watch [this video](https://www.youtube.com/watch?v=ok9mHOuvVSI).

After you have your _Cloudinary URL_ paste it into the .env file like so `CLOUDINARY_URL=cloudinary://API_KEY:API_SECRET@CLOUD_NAME`.

Once done retrieving the `DB_URL` and `CLOUDINARY_URL` we can start setting up the backend server.

### 5. Backend Server Setup

Before we start up the server, make sure you have the .env file created and setup. A template can be found at the bottom of the document.
The `PORT` can be set whatever port you want, just make sure to avoid ports that are already in use on your machine.
If you left the pgAdmin setup as defult, remember to avoid port 5432 as that will be the port that your PostgreSQL will be running at.
The `JWT_SECRET` is just a string used to create and validate the JWT tokens, set it to whatever string you like.
The `DB_URL`, `CLOUDINARY_URL` and `CLOUDINARY_UPLOAD_PRESET` should be gotten by following the instructions in the previous 2 sections.

Due to Golang's lightweight nature, the required modules are already included in the repository.
However, if you are still missing some modules, you can open up a terminal in the main folder
and run `go mod vendor` to install the required modules.

Now that all the initial stuff is out of the way, we can finally run the project. Make sure that your .env file is in the main directory and not in
a sub-directory. To build the project, open up a new terminal in the main project directory, run `go build`.
After building, you should now see an .exe file has been generated in the main directory.

![go build](https://github.com/Gareth2YuSheng/hwh-backend/blob/main/docs/readmeImages/go_build.png)

To run the .exe file, you can run `./hwh-backend.exe` (or whatever the exe file is called) in the terminal. When running the .exe file for the first time,
make sure to allow it access to the private network if not the server will not work.
If everything was setup correctly, you should see > `INFO - Server starting on port: 8080` in the terminal.

![go run](https://github.com/Gareth2YuSheng/hwh-backend/blob/main/docs/readmeImages/go_run.png)

If something goes wrong, do double check that
the .env file was setup correctly, alternatively you can check the terminal for any errors to try and troubleshoot what went wrong.

## .env File template

```
PORT=
DB_URL=postgresql://[user[:password]@][netloc][:port][/dbname][?param1=value1&...]
JWT_SECRET=
CLOUDINARY_URL=cloudinary://API_KEY:API_SECRET@CLOUD_NAME
CLOUDINARY_UPLOAD_PRESET=
```

for local instances, make sure to set the sslmode param to disable, I have a provided an example of my own local DB url
