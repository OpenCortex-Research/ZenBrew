use std::env;

mod repo;

fn main() {
    let args: Vec<String> = env::args().collect();

    let run_mode = &args[1];
    if run_mode == "--help" || run_mode == "-h" {
        println!("Usage: zen_brew <mode>");
        println!("Modes:");
        println!("    --help, -h: Print this help message");
        println!("    --version, -v: Print version information");
    } else if run_mode == "--version" || run_mode == "-v" {
        println!("zen_brew {}", env!("CARGO_PKG_VERSION"));
    } else if run_mode == "--install" || run_mode == "-i" {
        install_packages();
    } else if run_mode == "--remove" || run_mode == "-r" {
        remove_packages();
    } else if run_mode == "--update" || run_mode == "-u" {
        update_packages();
    }
    println!("{:?}", args);
}

fn install_packages() {
    println!("Install packages");
}

fn update_packages() {
    println!("Update packages");
}

fn remove_packages() {
    println!("Remove packages");
}
