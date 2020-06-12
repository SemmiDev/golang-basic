package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Employee struct {
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Gender         string `json:"gender"`
	Email          string `json:"email"`
	DateOfBirth    string `json:"date_of_birth"`
	CountryOfBirth string `json:"country_of_birth"`
}

func decodeJSON(data []byte) []Employee {
	var emp []Employee

	err := json.Unmarshal(data, &emp)
	if err != nil {
		log.Fatal(err)
	}
	return emp
}

func showToConsole() {

}

func main() {
	empJSON := `[{"first_name":"Scarlet","last_name":"Ipplett","gender":"Female","email":"sipplett0@yale.edu","date_of_birth":"2019-05-27","country_of_birth":"Nigeria"},
		{"first_name":"Kelsy","last_name":"Bolitho","gender":"Female","email":"kbolitho1@abc.net.au","date_of_birth":"2020-01-15","country_of_birth":"Mexico"},
		{"first_name":"Lisbeth","last_name":"Bywaters","gender":"Female","email":"lbywaters2@mediafire.com","date_of_birth":"2019-10-22","country_of_birth":"United States"},
		{"first_name":"Zane","last_name":"Holtum","gender":"Male","email":"zholtum3@usa.gov","date_of_birth":"2019-04-28","country_of_birth":"France"},
		{"first_name":"Lyon","last_name":"Haet","gender":"Male","email":"lhaet4@examiner.com","date_of_birth":"2020-02-09","country_of_birth":"Brazil"},
		{"first_name":"Keir","last_name":"Fotherby","gender":"Male","email":"kfotherby5@mtv.com","date_of_birth":"2020-03-28","country_of_birth":"Egypt"},
		{"first_name":"Tiphanie","last_name":"Dur","gender":"Female","email":"tdur6@google.ca","date_of_birth":"2019-07-12","country_of_birth":"China"},
		{"first_name":"Vivyan","last_name":"Pollen","gender":"Female","email":"vpollen7@soup.io","date_of_birth":"2019-10-05","country_of_birth":"Indonesia"},
		{"first_name":"Tessi","last_name":"Barnish","gender":"Female","email":"tbarnish8@symantec.com","date_of_birth":"2020-02-11","country_of_birth":"China"},
		{"first_name":"Shirlene","last_name":"Rentenbeck","gender":"Female","email":"srentenbeck9@twitpic.com","date_of_birth":"2020-01-24","country_of_birth":"Guatemala"},
		{"first_name":"Lyndsay","last_name":"Sarrell","gender":"Female","email":"lsarrella@walmart.com","date_of_birth":"2020-01-03","country_of_birth":"Indonesia"},
		{"first_name":"Lucias","last_name":"Ragless","gender":"Male","email":"lraglessb@google.com.hk","date_of_birth":"2020-02-14","country_of_birth":"South Africa"},
		{"first_name":"Urbano","last_name":"Seamons","gender":"Male","email":"useamonsc@cisco.com","date_of_birth":"2020-02-01","country_of_birth":"Thailand"},
		{"first_name":"Orrin","last_name":"Emblow","gender":"Male","email":"oemblowd@dmoz.org","date_of_birth":"2019-10-16","country_of_birth":"China"},
		{"first_name":"Kessia","last_name":"Bortoluzzi","gender":"Female","email":"kbortoluzzie@sogou.com","date_of_birth":"2020-02-23","country_of_birth":"Mexico"},
		{"first_name":"Giorgia","last_name":"Sprason","gender":"Female","email":"gsprasonf@wunderground.com","date_of_birth":"2019-06-21","country_of_birth":"Indonesia"},
		{"first_name":"Britt","last_name":"Wych","gender":"Female","email":"bwychg@slashdot.org","date_of_birth":"2020-03-10","country_of_birth":"Russia"},
		{"first_name":"Elvina","last_name":"Seebert","gender":"Female","email":"eseeberth@odnoklassniki.ru","date_of_birth":"2019-10-21","country_of_birth":"Philippines"},
		{"first_name":"Dianemarie","last_name":"Tynnan","gender":"Female","email":"dtynnani@vistaprint.com","date_of_birth":"2020-01-14","country_of_birth":"Afghanistan"},
		{"first_name":"Lucila","last_name":"Curd","gender":"Female","email":"lcurdj@tumblr.com","date_of_birth":"2019-11-23","country_of_birth":"China"},
		{"first_name":"Molli","last_name":"Newitt","gender":"Female","email":"mnewittk@mail.ru","date_of_birth":"2020-03-25","country_of_birth":"Afghanistan"},
		{"first_name":"Ross","last_name":"Treacy","gender":"Male","email":"rtreacyl@weather.com","date_of_birth":"2019-12-29","country_of_birth":"Czech Republic"},
		{"first_name":"Marv","last_name":"Cropton","gender":"Male","email":"mcroptonm@about.com","date_of_birth":"2019-09-14","country_of_birth":"Indonesia"},
		{"first_name":"Valeda","last_name":"Amer","gender":"Female","email":"vamern@1688.com","date_of_birth":"2020-01-29","country_of_birth":"China"},
		{"first_name":"Alfred","last_name":"Cinderey","gender":"Male","email":"acindereyo@people.com.cn","date_of_birth":"2019-06-04","country_of_birth":"Portugal"},
		{"first_name":"Lars","last_name":"Slamaker","gender":"Male","email":"lslamakerp@cdc.gov","date_of_birth":"2019-11-30","country_of_birth":"France"},
		{"first_name":"Dorette","last_name":"Fozzard","gender":"Female","email":"dfozzardq@vinaora.com","date_of_birth":"2019-10-25","country_of_birth":"Indonesia"},
		{"first_name":"Grenville","last_name":"Sturr","gender":"Male","email":"gsturrr@sourceforge.net","date_of_birth":"2019-06-22","country_of_birth":"Venezuela"},
		{"first_name":"Cookie","last_name":"Kettlestring","gender":"Female","email":"ckettlestrings@livejournal.com","date_of_birth":"2019-11-13","country_of_birth":"Iran"},
		{"first_name":"Ward","last_name":"Blaksland","gender":"Male","email":"wblakslandt@ycombinator.com","date_of_birth":"2020-01-17","country_of_birth":"China"},
		{"first_name":"Antoni","last_name":"Hadenton","gender":"Male","email":"ahadentonu@scribd.com","date_of_birth":"2019-12-05","country_of_birth":"Brazil"},
		{"first_name":"Graehme","last_name":"Riccio","gender":"Male","email":"gricciov@guardian.co.uk","date_of_birth":"2019-09-28","country_of_birth":"France"},
		{"first_name":"Sasha","last_name":"Mayne","gender":"Male","email":"smaynew@woothemes.com","date_of_birth":"2019-09-23","country_of_birth":"United States"},
		{"first_name":"Sal","last_name":"Balham","gender":"Female","email":"sbalhamx@whitehouse.gov","date_of_birth":"2019-06-12","country_of_birth":"Philippines"},
		{"first_name":"Marillin","last_name":"Mitchell","gender":"Female","email":"mmitchelly@bing.com","date_of_birth":"2020-02-01","country_of_birth":"Poland"},
		{"first_name":"Natale","last_name":"Bunten","gender":"Male","email":"nbuntenz@unesco.org","date_of_birth":"2019-07-21","country_of_birth":"China"},
		{"first_name":"Arty","last_name":"Worthington","gender":"Male","email":"aworthington10@sfgate.com","date_of_birth":"2019-11-30","country_of_birth":"China"},
		{"first_name":"Jeddy","last_name":"Thresher","gender":"Male","email":"jthresher11@sphinn.com","date_of_birth":"2020-02-03","country_of_birth":"United States"},
		{"first_name":"Bobbye","last_name":"Ciciotti","gender":"Female","email":"bciciotti12@disqus.com","date_of_birth":"2020-04-09","country_of_birth":"Indonesia"},
		{"first_name":"Bartholomew","last_name":"Sharland","gender":"Male","email":"bsharland13@newsvine.com","date_of_birth":"2019-10-01","country_of_birth":"Thailand"},
		{"first_name":"Emmalyn","last_name":"Vickar","gender":"Female","email":"evickar14@example.com","date_of_birth":"2019-05-25","country_of_birth":"China"},
		{"first_name":"Lazar","last_name":"Biggam","gender":"Male","email":"lbiggam15@fotki.com","date_of_birth":"2020-01-16","country_of_birth":"Canada"},
		{"first_name":"Lev","last_name":"Fawloe","gender":"Male","email":"lfawloe16@themeforest.net","date_of_birth":"2020-01-24","country_of_birth":"China"},
		{"first_name":"Rudolf","last_name":"Custy","gender":"Male","email":"rcusty17@nifty.com","date_of_birth":"2019-07-25","country_of_birth":"Albania"},
		{"first_name":"Leona","last_name":"Sibthorp","gender":"Female","email":"lsibthorp18@w3.org","date_of_birth":"2019-10-12","country_of_birth":"Czech Republic"},
		{"first_name":"Gilda","last_name":"Checklin","gender":"Female","email":"gchecklin19@last.fm","date_of_birth":"2019-07-16","country_of_birth":"Canada"},
		{"first_name":"Cinnamon","last_name":"Clitherow","gender":"Female","email":"cclitherow1a@cbc.ca","date_of_birth":"2019-04-20","country_of_birth":"Tajikistan"},
		{"first_name":"Jessie","last_name":"Manilo","gender":"Male","email":"jmanilo1b@virginia.edu","date_of_birth":"2019-08-14","country_of_birth":"Ireland"},
		{"first_name":"Paulette","last_name":"MacNess","gender":"Female","email":"pmacness1c@eepurl.com","date_of_birth":"2019-12-25","country_of_birth":"Indonesia"},
		{"first_name":"Nikki","last_name":"Heaseman","gender":"Male","email":"nheaseman1d@newsvine.com","date_of_birth":"2019-08-24","country_of_birth":"Indonesia"},
		{"first_name":"Chaunce","last_name":"Bradbeer","gender":"Male","email":"cbradbeer1e@usatoday.com","date_of_birth":"2020-03-01","country_of_birth":"Indonesia"},
		{"first_name":"Ernst","last_name":"Gilffillan","gender":"Male","email":"egilffillan1f@odnoklassniki.ru","date_of_birth":"2019-10-12","country_of_birth":"China"},
		{"first_name":"Brigitte","last_name":"Gloves","gender":"Female","email":"bgloves1g@ustream.tv","date_of_birth":"2019-07-06","country_of_birth":"China"},
		{"first_name":"Linet","last_name":"De Nisco","gender":"Female","email":"ldenisco1h@ning.com","date_of_birth":"2020-01-17","country_of_birth":"Denmark"},
		{"first_name":"Jedediah","last_name":"Allender","gender":"Male","email":"jallender1i@jalbum.net","date_of_birth":"2020-03-04","country_of_birth":"China"},
		{"first_name":"Minny","last_name":"Busst","gender":"Female","email":"mbusst1j@bloomberg.com","date_of_birth":"2019-04-30","country_of_birth":"Malaysia"},
		{"first_name":"Livvy","last_name":"Sutliff","gender":"Female","email":"lsutliff1k@about.me","date_of_birth":"2020-03-13","country_of_birth":"China"},
		{"first_name":"Dallas","last_name":"Steeden","gender":"Male","email":"dsteeden1l@surveymonkey.com","date_of_birth":"2019-09-17","country_of_birth":"Russia"},
		{"first_name":"Gayelord","last_name":"Briers","gender":"Male","email":"gbriers1m@yolasite.com","date_of_birth":"2019-07-15","country_of_birth":"France"},
		{"first_name":"Cedric","last_name":"Picard","gender":"Male","email":"cpicard1n@reuters.com","date_of_birth":"2019-05-05","country_of_birth":"China"},
		{"first_name":"Hector","last_name":"Mehaffey","gender":"Male","email":"hmehaffey1o@xinhuanet.com","date_of_birth":"2019-05-13","country_of_birth":"Mexico"},
		{"first_name":"Jerrilyn","last_name":"Allston","gender":"Female","email":"jallston1p@wisc.edu","date_of_birth":"2020-04-19","country_of_birth":"Colombia"},
		{"first_name":"Hedwiga","last_name":"Ellse","gender":"Female","email":"hellse1q@apple.com","date_of_birth":"2019-12-07","country_of_birth":"South Korea"},
		{"first_name":"Bruis","last_name":"Hoyle","gender":"Male","email":"bhoyle1r@netscape.com","date_of_birth":"2019-12-09","country_of_birth":"Israel"},
		{"first_name":"Shelagh","last_name":"MacGillivrie","gender":"Female","email":"smacgillivrie1s@uol.com.br","date_of_birth":"2019-06-10","country_of_birth":"Russia"},
		{"first_name":"Atlanta","last_name":"Backe","gender":"Female","email":"abacke1t@creativecommons.org","date_of_birth":"2019-06-01","country_of_birth":"Philippines"},
		{"first_name":"Damian","last_name":"Widd","gender":"Male","email":"dwidd1u@issuu.com","date_of_birth":"2019-06-14","country_of_birth":"Germany"},
		{"first_name":"Nananne","last_name":"Brimming","gender":"Female","email":"nbrimming1v@google.com","date_of_birth":"2019-12-13","country_of_birth":"Indonesia"},
		{"first_name":"Zebadiah","last_name":"Ianni","gender":"Male","email":"zianni1w@1688.com","date_of_birth":"2020-01-09","country_of_birth":"Indonesia"},
		{"first_name":"Archy","last_name":"Kuhle","gender":"Male","email":"akuhle1x@blogtalkradio.com","date_of_birth":"2019-12-18","country_of_birth":"Azerbaijan"},
		{"first_name":"Audy","last_name":"Doumer","gender":"Female","email":"adoumer1y@usda.gov","date_of_birth":"2020-01-08","country_of_birth":"Serbia"},
		{"first_name":"Mick","last_name":"Horley","gender":"Male","email":"mhorley1z@comsenz.com","date_of_birth":"2019-10-24","country_of_birth":"Austria"},
		{"first_name":"Edita","last_name":"Franzoli","gender":"Female","email":"efranzoli20@google.com.hk","date_of_birth":"2019-05-02","country_of_birth":"Indonesia"},
		{"first_name":"Gallagher","last_name":"Feldman","gender":"Male","email":"gfeldman21@usa.gov","date_of_birth":"2019-06-23","country_of_birth":"Colombia"},
		{"first_name":"Paddy","last_name":"Richemont","gender":"Male","email":"prichemont22@java.com","date_of_birth":"2019-10-07","country_of_birth":"Dominican Republic"},
		{"first_name":"Gibbie","last_name":"Kissock","gender":"Male","email":"gkissock23@bluehost.com","date_of_birth":"2019-10-19","country_of_birth":"Ukraine"},
		{"first_name":"Thadeus","last_name":"Ricart","gender":"Male","email":"tricart24@simplemachines.org","date_of_birth":"2019-12-08","country_of_birth":"Vietnam"},
		{"first_name":"Morissa","last_name":"Nouch","gender":"Female","email":"mnouch25@fema.gov","date_of_birth":"2019-10-20","country_of_birth":"Thailand"},
		{"first_name":"Caryl","last_name":"Fowle","gender":"Male","email":"cfowle26@nytimes.com","date_of_birth":"2019-04-27","country_of_birth":"Vietnam"},
		{"first_name":"Carey","last_name":"Roycraft","gender":"Male","email":"croycraft27@gnu.org","date_of_birth":"2019-12-21","country_of_birth":"Ireland"},
		{"first_name":"Billi","last_name":"Mantha","gender":"Female","email":"bmantha28@stanford.edu","date_of_birth":"2019-05-18","country_of_birth":"Pakistan"},
		{"first_name":"Dania","last_name":"Bicknell","gender":"Female","email":"dbicknell29@soundcloud.com","date_of_birth":"2019-08-26","country_of_birth":"China"},
		{"first_name":"Banky","last_name":"Malam","gender":"Male","email":"bmalam2a@zdnet.com","date_of_birth":"2020-01-30","country_of_birth":"Morocco"},
		{"first_name":"Meagan","last_name":"Linney","gender":"Female","email":"mlinney2b@seesaa.net","date_of_birth":"2019-10-02","country_of_birth":"China"},
		{"first_name":"Charlie","last_name":"Duck","gender":"Male","email":"cduck2c@goodreads.com","date_of_birth":"2019-07-20","country_of_birth":"Indonesia"},
		{"first_name":"Marven","last_name":"Spurr","gender":"Male","email":"mspurr2d@opensource.org","date_of_birth":"2019-06-25","country_of_birth":"Solomon Islands"},
		{"first_name":"Gamaliel","last_name":"Campsall","gender":"Male","email":"gcampsall2e@deliciousdays.com","date_of_birth":"2019-08-22","country_of_birth":"Jordan"},
		{"first_name":"Guss","last_name":"Melburg","gender":"Male","email":"gmelburg2f@disqus.com","date_of_birth":"2020-01-11","country_of_birth":"Denmark"},
		{"first_name":"Konstantine","last_name":"Reame","gender":"Male","email":"kreame2g@usnews.com","date_of_birth":"2019-06-01","country_of_birth":"United States"},
		{"first_name":"Anjela","last_name":"Spendlove","gender":"Female","email":"aspendlove2h@csmonitor.com","date_of_birth":"2020-02-05","country_of_birth":"China"},
		{"first_name":"Alden","last_name":"Fulun","gender":"Male","email":"afulun2i@feedburner.com","date_of_birth":"2019-06-11","country_of_birth":"Albania"},
		{"first_name":"Boycie","last_name":"Jeannin","gender":"Male","email":"bjeannin2j@biblegateway.com","date_of_birth":"2019-11-16","country_of_birth":"Indonesia"},
		{"first_name":"Addy","last_name":"Wickwarth","gender":"Male","email":"awickwarth2k@xinhuanet.com","date_of_birth":"2019-11-12","country_of_birth":"Canada"},
		{"first_name":"Frances","last_name":"Arkcoll","gender":"Female","email":"farkcoll2l@ameblo.jp","date_of_birth":"2019-09-13","country_of_birth":"China"},
		{"first_name":"Carrissa","last_name":"Petow","gender":"Female","email":"cpetow2m@shutterfly.com","date_of_birth":"2019-12-09","country_of_birth":"China"},
		{"first_name":"Christiana","last_name":"Lannin","gender":"Female","email":"clannin2n@g.co","date_of_birth":"2019-10-02","country_of_birth":"Indonesia"},
		{"first_name":"Dacy","last_name":"Faires","gender":"Female","email":"dfaires2o@google.com.br","date_of_birth":"2019-11-08","country_of_birth":"China"},
		{"first_name":"Tressa","last_name":"Thorouggood","gender":"Female","email":"tthorouggood2p@t-online.de","date_of_birth":"2019-09-06","country_of_birth":"China"},
		{"first_name":"Yuri","last_name":"Gildea","gender":"Male","email":"ygildea2q@networksolutions.com","date_of_birth":"2020-01-13","country_of_birth":"Indonesia"},
		{"first_name":"Pat","last_name":"Babbs","gender":"Male","email":"pbabbs2r@ucsd.edu","date_of_birth":"2019-12-01","country_of_birth":"China"}]`
	dataJSON := decodeJSON([]byte(empJSON))

	for _, v := range dataJSON {
		fmt.Println(" First Name  : ", v.FirstName)
		fmt.Println(" Last Name   : ", v.LastName)
		fmt.Println(" Gender    : ", v.Gender)
		fmt.Println(" Email    : ", v.Email)
		fmt.Println(" Date Of Birth    : ", v.DateOfBirth)
		fmt.Println(" Country Of Birth    : ", v.CountryOfBirth+"\n\n")
	}

	c, err := os.Create("JSONencodingtoTxt.txt")
	if err != nil {
		log.Fatal(err)
	}

	var index int = 1
	for _, v := range dataJSON {
		nilaiString := strconv.Itoa(index)

		c.WriteString(" No : " + nilaiString)
		c.WriteString(" First Name  : " + v.FirstName)
		c.WriteString(" Last Name   : " + v.LastName)
		c.WriteString(" Gender      : " + v.Gender)
		c.WriteString(" Email       : " + v.Email)
		c.WriteString(" Date Of Birth    : " + v.DateOfBirth)
		c.WriteString(" COuntry Of Birt  : " + v.CountryOfBirth + "\n\n")
		index++
	}
	defer c.Close()
	fmt.Println("OK!")
}
