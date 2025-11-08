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
| `--format` | Choose outputting format ( flags; c = checksum, p = path, d = date) | `--format cdp` |

---

### Examples

#### Generate SHA256 checksums for all files in a folder
```bash
checksum --folder ./myfiles
```
#### Generate SHA256 checksum for one file
```bash
checksum --file ./test.txt
```
#### Generate SHA256 checksum for one file, but output only the path and date
```bash
checksum --file ./test.txt --format pd
```

## Help Message
```bash
Usage:
  checksum [--file <path>] [--folder <path>]

Flags:
  --file     Path to a single file
  --folder   Path to a folder for batch checksums
  --format   Choose output format flags: `c` = checksum, `p` = path, `d` = date (default = `cpd`)

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
8. Tein tahallaan virheen ja reverttasin
https://i.imgur.com/nuDqV5m.png
9. Yritin cherry-pick:ata commitin mainista, mutta ei tee mitään koska on jo devissä ( git cherry-pick 39e2854 )
10. Lisäsin tägin v1.0.0 joka on stable versio
11. Loin .gitignore tiedoston, ja testauksen vuoksi jätin binäärin.
12. Käytin rebase komentoa, ja kaikki oli jo synkronoituna.
13. Korjasin dokumentoinnin.
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


704  git status
705  git stash push -u -m "puolvälissä"
706  git status
707  git stash list
708  git stash pop
731  go build main.go; ./main --file README.md --format cpd
737  git add .
738  git commit -m "added --format feature"
739  git checkout dev
740  it fetch origin
741  clear
742  git fetch origin
743  git rebase main
744  git checkout main
745  git merge dev
746  git push origin main && git push origin dev
747  clear
748  git checkout dev
749  git rebase main
750  clear
751  echo "tahallinen virhe" >> main.go
752  git add main.go
753  git commit -m "purposeful error added"
754  git log --oneline -1
755  git revert HEAD
756  git log --oneline -2
757  clear
758  git checkout main && git log --oneline
759  git checkout dev
760  git cherry-pick 39e2854
761  git log --oneline -5
762  git checkout main
763  clear
764  git merge dev
765  git tag v1.0.0 && git push origin v1.0.0
766  git checkout dev
767  clear
768  git checkout dev
769  nano .gitignore
770  git add .gitignore && git commit -m "added .gitignore for binaries"
774  ls
775  go build main.go; ./main --file README.md --format cdd
776  clear
777  git checkout main && git merge dev && git push origin main
778  git checkout dev
779  git fetch origin
780  git rebase main
781  ls
782  git stash push -u -m "ei valmista kaikki"
783  git stash pop
784  git add README.md && git commit -m "updated readme.md"
785  git rebase main
786  clear
```

---