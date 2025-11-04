# checksum – SHA256 File Checksum Generator

A simple command-line checksum generator written in **Golang**.  
It generates **SHA256 checksums** for single files or entire folders.  
This tool can help you verify the authenticity of your files.

---

## Usage

checksum [FLAGS]

### ⚙ Flags

| Flag | Description | Example |
|------|--------------|----------|
| `--file` | Generate checksum for a single file | `--file example.txt` |
| `--folder` | Generate checksums for all files in a folder | `--folder ./downloads` |

---

### Examples

#### Generate SHA256 checksums for all files in a folder
```bash
checksum --folder ./myfiles
```
#### Generate SHA256 checksum for one file
```
checksum --file ./test.txt
```

## Help Message
```bash
Usage:
  checksum [--file <path>] [--folder <path>]

Flags:
  --file     Path to a single file
  --folder   Path to a folder for batch checksums
```

---

# Opinnäyte kohdat

Alla ovat vaiheet ja Git-komennot, joita käytin harjoituksessa versionhallinnan, haarojen ja commitien oppimisen osana.

### Mitä tein

1. Loin private ja public SSH avaimet, lisäsin ne käyttäjälle ja kloonasin repon SSH kanssa ( `git clone git@github.com:HuzoDumpster/checksums.git` ).
2. Loin toisen "dev" haaran jossa loin V1 sovelluksesta.

---

### CLI History

```
549  ssh-keygen -t ed25519 -C "sc16282@student.samk.fi"
557  ls /home/unknown/.ssh
558  mv * /home/unknown/.ssh
561  ssh -T git@github.com
563  git clone git@github.com:HuzoDumpster/checksum.git
564  cd checksum
565  nano README.md
570  git add README.md
571  git commit -m "Alkuperäisen README.md lisäys."
573  git checkout -b dev
574  git branch -a

```

---