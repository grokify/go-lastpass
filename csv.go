package lastpass

import "github.com/grokify/gocharts/v2/data/table"

type Accounts []Account

func CSVColumns() []string {
	return []string{
		"url",
		"username",
		"password",
		"extra",
		"name",
		"grouping",
		"fav",
		"totp",
	}
}

func (acs Accounts) Table() table.Table {
	tbl := table.NewTable("LastPass accounts")
	tbl.Columns = CSVColumns()
	for _, ac := range acs {
		if (Account{}) == ac {
			continue
		}
		row := acs.TableRow(ac)
		if len(row) > 0 {
			tbl.Rows = append(tbl.Rows, row)
		}
	}
	return tbl
}

func (acs Accounts) TableRow(a Account) []string {
	if (Account{}) == a {
		return []string{}
	}
	return []string{
		a.URL,
		a.Username,
		a.Password,
		a.Extra,
		a.Name,
		a.Grouping,
		a.Fav,
		a.TOTP,
	}
}

func (acs Accounts) WriteCSV(filename string) error {
	tbl := acs.Table()
	return tbl.WriteCSV(filename)
}

type Account struct {
	// url – typically the login URL, but you can use http://sn to create a Secure Note
	URL string `json:"url"`
	// username – the username used for the site entry
	Username string `json:"username"`
	// password – the password used for the site entry
	Password string `json:"password"`
	// extra – can be either the "notes" section of a site entry or the body of a secure note
	Extra string `json:"extra"`
	// name – the name to identify the site entry or secure note in your vault
	Name string `json:"name"`
	// grouping – the folder where you would like the item to be stored in your vault
	Grouping string `json:"grouping"`
	// fav – enter a 1, if the entry should be added as a Favorite but by default the value will be 0 (indicating that it should not be added as a Favorite)
	Fav string `json:"fav"`
	// totp – enter the TOTP code generated for your site entry (learn more about this feature)
	TOTP string `json:"totp"`
}
