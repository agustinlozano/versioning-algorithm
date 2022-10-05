use std::fs::File;
use std::io::{BufReader, ErrorKind};
use std::time::Instant;
use std::io::Read;

fn main() {
    let path = "../IFCs/myIFC.ifc";

    // Abrir el archivo
    let file = File::open(path)
        .expect("Failed to open file");

    let mut buf = BufReader::new(file);

    let mut x = 0u8;

    let mut num_bytes = 0usize;

    let t1 = Instant::now();

    let mut bytes = [0; 512];

    loop {
        match buf.read(&mut bytes) {
            Ok(0) => break,
            Ok(n) => {
                for i in 0..n {
                    num_bytes += 1;
                    x += bytes[i];
                }
            }
            Err(ref e) if e.kind() == ErrorKind::Interrupted => continue,
            Err(e) => panic!("{:?}", e),
        };
    }

    let dur = t1.elapsed().as_secs_f64();
    let num_bytes = (num_bytes as f64) / 1_048_576f64;

    println!("RESULT: {}", x);

    println!("Read speed: {:.1} MiB / s for {}", num_bytes / dur, num_bytes);

}