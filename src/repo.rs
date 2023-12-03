use std::{fs::File, io::Error};
use serde_json;
use serde::Deserialize;
use reqwest;

#[derive(Debug, Deserialize)]
struct RepoLink {
    name: String,
    url: String,
    sum: String,
}


// Load the repo file from disk
pub fn import_local_repo(file_location: Option<String>) {
    println!("Importing local repo");
    let file = File::open(file_location.unwrap_or("/media/p1/OpenCortex/ZenBrew/local_repo.json".to_string())).unwrap();
    let json: Vec<RepoLink> = serde_json::from_reader(file).expect("JSON was not well-formatted");
}

#[derive(Debug, Deserialize)]
pub struct Package {
    name: String,
    versions: Vec<PackageVersion>,
    author: String,
    description: String,
    sha256_sum: String,
}

impl Package {

    // Create a package from it's defining JSON file
    pub fn create(file: String) -> Result<Package, Error> {
        Err("Error")
    }
}

#[derive(Debug, Deserialize)]
struct PackageVersion {
    version: String,
    branch: String,
    install_url: String,
    update_url: String,
    remove_url: String,
}

#[derive(Debug, Deserialize)]
pub struct Repo {
    name: String,
    url: String,
    author: String,
    sha256_sum: String,
    packages: Vec<RepoLink>,
}

impl Repo {

    // Download repo descriptor and create repo object
    pub async fn new(name: String, url: String, sum: String) -> Result<Repo, Error> {

        // Download file
        let response = reqwest::get(&url).await.expect("error");

        // Turn response to json
        let repo_file: Repo = serde_json::from_str(&response.text().await.expect("error")).expect("JSON was not well-formatted");

        // Check repo contains the correct infomation
        if repo_file.name == name || repo_file.sha256_sum == sum || repo_file.url == url {
            println!("Repo does pass checks");
        } else {
            println!("Repo does not pass checks");
        }

        // Return new repo object
        Ok(Repo{
            name, url, author: repo_file.author, sha256_sum: sum, packages: repo_file.packages
        })
    }

    pub async fn download_package(mut self, name: String, version: String) -> Result<Package, Error> {
        for package in self.packages {
            if package.name == name {
                let response = reqwest::get(package.url).await?;
                let mut file = match File::create(("/media/p1/OpenCortex/ZenBrew/packages/{:?}.json", &name)) {
                    Err(why) => panic!("couldn't create {}", why),
                    Ok(file) => file,
                };
                file.write_all(&response.bytes().await?)
                    .expect("Unable to write to file");
                file.close();
                let package_out = Package::new(("/media/p1/OpenCortex/ZenBrew/packages/{:?}.json", name.clone()));
                Ok(package_out)
            }
        }
        Err("NO")
    }
}