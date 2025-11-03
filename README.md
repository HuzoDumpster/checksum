# checksum – Multi-file Checksum Generator

A simple command-line checksum generator written in **Golang**.
It can generate checksums for single files or entire folders using multiple algorithms, with flexible output options.
This tool may help you in verifying the authentcity of the files.

---

## Usage

checksum [FLAGS]


### ⚙Flags

| Flag | Description | Example |
|------|--------------|----------|
| `--file` | Generate checksum for a single file | `--file example.txt` |
| `--folder` | Generate checksums for all files in a folder | `--folder ./downloads` |
| `--algo` | Set the checksum algorithm (`sha256`, `md5`, etc.) | `--algo sha256` |
| `--out` | Specify output directory for checksum files | `--out ./output/` |
| `--format` | Choose output format flags: `c` = checksum, `p` = path, `d` = date (default = `cpd`) | `--format cp` |
| `--difference` | Compare output with a this reference checksum file. This reference file must be done with this tool or have similar format. | `--difference ./referenxce_file.txt` |

---

### Examples


# Generate SHA256 checksums for all files in a folder
`checksum --folder ./myfiles --algo sha256`

# Generate MD5 checksum for one file and output to a folder
`checksum --file ./test.txt --algo md5 --out ./sums`

# Custom format (only checksum + path)
`checksum --folder ./data --algo sha512 --format cp`


---

## Help Message

```
Usage:
  checksum [--file <path>] [--folder <path>] [--algo <name>] [--out <dir>] [--format <flags>]

Flags:
  --file     Path to a single file
  --folder   Path to a folder for batch checksums
  --algo     Checksum algorithm (md5, sha1, sha256, sha512)
  --out      Output directory for result file
  --format   Output format: c (checksum), p (path), d (date). Default = cpd
```

---

# Opinnäyte kohdat

Alla ovat vaiheet ja Git-komennot, joita käytin harjoituksessa versionhallinnan, haarojen ja commitien oppimisen osana.

### Mitä tein

1. Loin private ja public SSH avaimet, lisäsin ne käyttäjälle ja kloonasin repon SSH kanssa ( `git clone git@github.com:HuzoDumpster/checksums.git` ).
2. Lisäsin `.gitignore`-tiedoston ja koodasin Go-tiedostot.
- Toteutin CLI-lippujen käsittelyn  
- Rakensin tarkistussummien generoinnin Go:n crypto-kirjastojen avulla  
- Harjoittelin Git-operaatioita (commit, branch, merge, push jne.)  
- Dokumentoin työvaiheet tähän README-tiedostoon  

---

### 💻 CLI History

> Paste your terminal or Git CLI history here, for example:

[REPLACE]
git init
git add .
git commit -m "Initial commit with Go checksum project"
git branch feature-checksum
git checkout feature-checksum
# ...
[REPLACE]

---

## 🧱 Technologies Used

- **Go (Golang)** — pääohjelmointikieli  
- **Git & GitHub** — versionhallinta ja repositorion ylläpito  
- **Markdown** — dokumentointiin  

---

## 🏁 Future Improvements

- Lisää rinnakkaisen tarkistussummien generoinnin tuki suurille kansioille  
- Mahdollisuus määritellä mukautetut tulostiedoston nimet  
- Lisää tarkistusmoodi (`--verify`) tarkistussummien vertailuun  

---

## 📄 License

MIT License — vapaasti käytettävä, muokattava ja opittavaksi.

---

### 👨‍💻 Author

Created as part of a school Git exercise by *[Your Name]*.
[REPLACE]

