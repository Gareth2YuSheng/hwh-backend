# Homework-Help

CVWO Assignment Project
Server written in GOLANG, Database using POSTGRESQL, Media storying using CLOUDINARY
By Gareth Too Yu Sheng

## Setting Up the Project

### 1. Database Setup

If you already have a PostgreSQL database instance setup, feel free to skip this section and use your own database url.

To setup a Local Instance of PostgreSQL, download pgAdmin [here](https://www.pgadmin.org/download/).
Select the specific one for your machine and install, make sure to also install PostgreSQL as well.
If you do not have it installed, it can be installed together with pgAdmin via the pgAdmin installer.
Make sure to set up a root password and you can leave the port number as the default.
Go ahead and leave the rest of the options as default as well and wait for the installation to finish.

Once installed, start up pgAdmin. Click on servers on the left hand side menu and enter the root password that you provided.
Under Servers > PostgreSQL > Databases, right click on Databases, click on Create > Database... to create an new database.
Name the database whatever you like, no SQL file is required as the server will handle all the table setup once connected to the database.

After setting up the database, create a .env file in the main project directory and enter your database url into the DB_URL variable.
A template for the env file required for this project is provided at the bottom of the document.

For a clearer installation and setup process of pgAdmin, you can watch [this video](https://www.youtube.com/watch?v=4qH-7w5LZsA).

### 2. Cloudinary Setup

A Cloudinary account will be required for media storage, not to worry as there is a free tier available.
Sign up for a Cloudinary account [here](https://cloudinary.com/users/register_free).

After signing up for an account, we need to note down your Cloudinary Url for the .env file.
First, login using your new account and go to the bottom left there will be a settings icon above your profile. Click on it to go to the settings page.
At the top of the left side menu, you will see your `CLOUD_NAME`, under _Product environment settings_ navigate to _Upload Presets_.
There we will need to add a new _Upload Preset_, click on \_Add Upload Preset\* on the top right and name it whatever you like, e.g. hw-help-preset,
give a name for your _Asset Folder_ e.g. hw-help-images and make sure to change the _Signing mode_ to **Unsigned**.
Change the _Generated public ID_ to "Use the filename of the uploaded file as the public ID" and leave the rest of the settings as the default.
Once done click Save on the top right.

Next we need the API keys, similarly under Product environment settings, navigate to API Keys.
Your Cloudinary Url looks like `cloudinary://API_KEY:API_SECRET@CLOUD_NAME`. There will already be 1 API Key created for you by default, it is
up to you whether you wish to use it or create a new one. Select an `API Key` of your choice and copy the `API Key` and `API Secret`
into your Cloudinary Url in the .env file. To copy the `API Secret` you need to click on the eye icon to view it first.
On this page at the top also provides you with a sample Cloudinary Url with your `CLOUD_NAME`.
Alternatively you can also copy this into the .env file, but you are will required to copy the `API Key` and `API Secret` manually.
If you are having trouble finding your `API Key`, `API Secret` and `Cloud Name`, you can watch [this video](https://www.youtube.com/watch?v=ok9mHOuvVSI).

Once done retrieving the `DB_URL` and `CLOUDINARY_URL` we can start setting up the backend server.

### 3. Backend Server Setup

Before we start up the server, make sure you have the .env file created and setup. A template can be found at the bottom of the document.
The `PORT` can be set whatever port you want, just make sure to avoid ports that are already in use on your machine.
If you left the pgAdmin setup as defult, remember to avoid port 5432 as that will be the port that your PostgreSQL will be running at.
The `JWT_SECRET` is just a string used to create and validate the JWT tokens, set it to whatever string you like.
The `DB_URL` and `CLOUDINARY_URL` should be gotten from following the instructions in the previous 2 sections.

After creating and setting up the .env file, we can now run the project. Due to Golang's lightweight nature, the required modules are already
included in the repository. However, if you are still missing some modules, you can open up a terminal in the main folder
and run `go mod vendor` to install the required modules.

Now that all the initial stuff is out of the way, we can finally run the project. Make sure that your .env file is in the main directory and not in
a sub-directory. To build the project, run `go build`. After building, you should now see an .exe file has been generated in the main directory.
To run the .exe file, you can run `./hwh-backend.exe` (or whatever the exe file is called) in the terminal. When running the .exe file for the first time,
make sure to allow it access to the private network if not the server will not work.
If everything was setup correctly, you should see > `INFO - Server starting on port: 8080` in the terminal. If something goes wrong, do double check that
the .env file was setup correctly, alternatively you can check the terminal for any errors to try and troubleshoot what went wrong.

## .env File template

```
PORT=8080
DB_URL=postgresql://[user[:password]@][netloc][:port][/dbname][?param1=value1&...] e.g. postgres://postgres:root@localhost:5432/hwhelp?sslmode=disable
JWT_SECRET=JWTSECRET
CLOUDINARY_URL=cloudinary://API_KEY:API_SECRET@CLOUD_NAME
```

for local instances, make sure to set the sslmode param to disable, I have a provided an example of my own local DB url
