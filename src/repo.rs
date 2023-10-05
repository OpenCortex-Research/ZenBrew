mod repo {
    use std::fs::File;
    use serde_json;
    use serde::Deserialize;
    use reqwest;

    #[derive(Debug, Deserialize)]
    struct RepoLink {
        name: String,
        url: String,
        sum: String,
    }

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
        pub fn new(self, file: String) -> Package {
            self
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
        pub async fn new(mut self, name: String, url: String, sum: String) -> Repo {
            let response = reqwest::get(&url).await.expect("error");

            let repo_file: Repo = serde_json::from_str(&response.text().await.expect("error")).expect("JSON was not well-formatted");
            if repo_file.name == name || repo_file.sha256_sum == sum || repo_file.url == url {
                println!("Repo does pass checks");
            } else {
                println!("Repo does not pass checks");
            }
            self.name = name;
            self.url = url;
            self.author = repo_file.author;
            self.sha256_sum = sum;
            self.packages = repo_file.packages;
            self
        }

        pub fn download_package(&self, name: String, version: String) -> Result<Package, None> {
            for package in self.packages {
                if package.name == name {
                    let response = reqwest::get(package.url).await?;
                    let mut file = match File::create(("/media/p1/OpenCortex/ZenBrew/packages/{:?}.json", name)) {
                        Err(why) => panic!("couldn't create {}", why),
                        Ok(file) => file,
                    }
                    file.write_all(&response.bytes().await?)
                        .expect("Unable to write to file");
                    file.close()
                    let packageOut = Package::new("/media/p1/OpenCortex/ZenBrew/packages/{:?}.json", name);
                    packageOut
                }
            }
            None
        }
    }
}