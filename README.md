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
3. Lisäsin README.md tekstipäivitys commitin, vaihdoin main/ haaraan ja mergasin sen devin kanssa, joten devin tiedostot kopioitui main/ haaraan.
4. Lisäsin main haaraan pää .go tiedoston ja pushasin päivitetyn main/ haaran, sekä dev/ haaran.
5. Lisäsin riville 56 tulostustekstiä molempiin harroihin hiema eri tekstillä ja sain merge conflictin 
```
┌──(unknown㉿Linux)-[~/Documents/Temporary/Versiohallinta/checksum]
└─$ git merge dev
Auto-merging main.go
CONFLICT (add/add): Merge conflict in main.go
Automatic merge failed; fix conflicts and then commit the result.
```
Korjasin tämän pitämällä dev haaran version teidostosta.
6. Aloin koodaamaan main.go tiedoston uutta ominaisuutta dev/ haarassa ja halusin vaihtaa nopeasti main/ haaraan, joten suoritin komennon `git stash` ja vaihdoin haaraa, sitten vaihdoin takaisin dev/ haaraan ja suoritin komennon `git stash pop`
https://i.imgur.com/XFS1REa.png
7. Koodasin loppuun uuden ominaisuuden ja loin commitin.
8. 
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
( Coded / Edited Files)
639  git status
640  git add README.md
641  git commit -m "Päivitetty README.md sovelluksen ensimmäiseen versioon."
642  git checkout main
643  git merge dev
644  git add main.go
645  git commit -m "Lisätty versio yksi (main.go)"
646  git push origin main
647  git push origin dev
```

---