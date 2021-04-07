package Gsensor

type SubDomainRespsonse struct {
	Pageprops struct {
		Apexdomain     string `json:"apexDomain"`
		Apexdomaindata struct {
			Success    bool   `json:"success"`
			Status     int    `json:"status"`
			Statustext string `json:"statusText"`
			Error      string `json:"error"`
			Data       struct {
				Meta struct {
					LimitReached bool `json:"limit_reached"`
					MaxPage      int  `json:"max_page"`
					Page         int  `json:"page"`
					TotalPages   int  `json:"total_pages"`
				} `json:"meta"`
				Records []struct {
					HostProvider []string `json:"host_provider"`
					Hostname     string   `json:"hostname"`
					MailProvider []string `json:"mail_provider"`
					OpenPageRank int      `json:"open_page_rank"`
				} `json:"records"`
				Total    int `json:"total"`
				Previews []struct {
					Hostname string `json:"hostname"`
					Rank     string `json:"rank"`
					Provider string `json:"provider"`
				} `json:"previews"`
			} `json:"data"`
			Asnrisklevel string `json:"asnRiskLevel"`
		} `json:"apexDomainData"`
		Dnsdata struct {
			Success    bool   `json:"success"`
			Status     int    `json:"status"`
			Statustext string `json:"statusText"`
			Error      string `json:"error"`
			Data       struct {
				AlexaRank  interface{} `json:"alexa_rank"`
				ApexDomain string      `json:"apex_domain"`
				CurrentDNS struct {
					A struct {
						FirstSeen string `json:"first_seen"`
						Values    []struct {
							H              interface{} `json:"h"`
							IP             string      `json:"ip"`
							IPCount        int         `json:"ip_count"`
							IPOrganization string      `json:"ip_organization"`
						} `json:"values"`
					} `json:"a"`
					Aaaa struct {
					} `json:"aaaa"`
					Mx struct {
						FirstSeen string `json:"first_seen"`
						Values    []struct {
							Hostname             string `json:"hostname"`
							HostnameCount        int    `json:"hostname_count"`
							HostnameOrganization string `json:"hostname_organization"`
							Priority             int    `json:"priority"`
						} `json:"values"`
					} `json:"mx"`
					Ns struct {
						FirstSeen string `json:"first_seen"`
						Values    []struct {
							Nameserver             string `json:"nameserver"`
							NameserverCount        int    `json:"nameserver_count"`
							NameserverOrganization string `json:"nameserver_organization"`
						} `json:"values"`
					} `json:"ns"`
					Soa struct {
						FirstSeen string `json:"first_seen"`
						Values    []struct {
							Email      string `json:"email"`
							EmailCount int    `json:"email_count"`
							TTL        int    `json:"ttl"`
						} `json:"values"`
					} `json:"soa"`
					Txt struct {
						FirstSeen string `json:"first_seen"`
						Values    []struct {
							Value string `json:"value"`
						} `json:"values"`
					} `json:"txt"`
				} `json:"current_dns"`
				Hostname       string `json:"hostname"`
				SubdomainCount int    `json:"subdomain_count"`
				History        struct {
					A struct {
						Data struct {
							Records []struct {
								Firstseen                 string   `json:"firstSeen"`
								Lastseen                  string   `json:"lastSeen"`
								Firstseenduration         string   `json:"firstSeenDuration"`
								Lastseenduration          string   `json:"lastSeenDuration"`
								Firstseenlastseenduration string   `json:"firstSeenLastSeenDuration"`
								Organizations             []string `json:"organizations"`
								Values                    []struct {
									IP string `json:"ip"`
								} `json:"values"`
							} `json:"records"`
							Previews []struct {
								Firstseen                 string   `json:"firstSeen"`
								Lastseen                  string   `json:"lastSeen"`
								Firstseenduration         string   `json:"firstSeenDuration"`
								Lastseenduration          string   `json:"lastSeenDuration"`
								Firstseenlastseenduration string   `json:"firstSeenLastSeenDuration"`
								Organizations             []string `json:"organizations"`
								Values                    []struct {
									IP string `json:"ip"`
								} `json:"values"`
							} `json:"previews"`
						} `json:"data"`
					} `json:"a"`
					Aaaa struct {
						Data struct {
							Records  []interface{} `json:"records"`
							Previews []interface{} `json:"previews"`
						} `json:"data"`
					} `json:"aaaa"`
					Mx struct {
						Data struct {
							Records []struct {
								Firstseen                 string   `json:"firstSeen"`
								Lastseen                  string   `json:"lastSeen"`
								Firstseenduration         string   `json:"firstSeenDuration"`
								Lastseenduration          string   `json:"lastSeenDuration"`
								Firstseenlastseenduration string   `json:"firstSeenLastSeenDuration"`
								Organizations             []string `json:"organizations"`
								Values                    []struct {
									Host string `json:"host"`
								} `json:"values"`
							} `json:"records"`
							Previews []struct {
								Firstseen                 string   `json:"firstSeen"`
								Lastseen                  string   `json:"lastSeen"`
								Firstseenduration         string   `json:"firstSeenDuration"`
								Lastseenduration          string   `json:"lastSeenDuration"`
								Firstseenlastseenduration string   `json:"firstSeenLastSeenDuration"`
								Organizations             []string `json:"organizations"`
								Values                    []struct {
									Host string `json:"host"`
								} `json:"values"`
							} `json:"previews"`
						} `json:"data"`
					} `json:"mx"`
					Ns struct {
						Data struct {
							Records []struct {
								Firstseen                 string   `json:"firstSeen"`
								Lastseen                  string   `json:"lastSeen"`
								Firstseenduration         string   `json:"firstSeenDuration"`
								Lastseenduration          string   `json:"lastSeenDuration"`
								Firstseenlastseenduration string   `json:"firstSeenLastSeenDuration"`
								Organizations             []string `json:"organizations"`
								Values                    []struct {
									Nameserver string `json:"nameserver"`
								} `json:"values"`
							} `json:"records"`
							Previews []struct {
								Firstseen                 string   `json:"firstSeen"`
								Lastseen                  string   `json:"lastSeen"`
								Firstseenduration         string   `json:"firstSeenDuration"`
								Lastseenduration          string   `json:"lastSeenDuration"`
								Firstseenlastseenduration string   `json:"firstSeenLastSeenDuration"`
								Organizations             []string `json:"organizations"`
								Values                    []struct {
									Nameserver string `json:"nameserver"`
								} `json:"values"`
							} `json:"previews"`
						} `json:"data"`
					} `json:"ns"`
					Soa struct {
						Data struct {
							Records []struct {
								Firstseen                 string        `json:"firstSeen"`
								Lastseen                  string        `json:"lastSeen"`
								Firstseenduration         string        `json:"firstSeenDuration"`
								Lastseenduration          string        `json:"lastSeenDuration"`
								Firstseenlastseenduration string        `json:"firstSeenLastSeenDuration"`
								Organizations             []interface{} `json:"organizations"`
								Values                    []struct {
									Email      string `json:"email"`
									EmailCount int    `json:"email_count"`
									TTL        int    `json:"ttl"`
								} `json:"values"`
							} `json:"records"`
							Previews []struct {
								Firstseen                 string        `json:"firstSeen"`
								Lastseen                  string        `json:"lastSeen"`
								Firstseenduration         string        `json:"firstSeenDuration"`
								Lastseenduration          string        `json:"lastSeenDuration"`
								Firstseenlastseenduration string        `json:"firstSeenLastSeenDuration"`
								Organizations             []interface{} `json:"organizations"`
								Values                    []struct {
									Email      string `json:"email"`
									EmailCount int    `json:"email_count"`
									TTL        int    `json:"ttl"`
								} `json:"values"`
							} `json:"previews"`
						} `json:"data"`
					} `json:"soa"`
					Txt struct {
						Data struct {
							Records []struct {
								Firstseen                 string        `json:"firstSeen"`
								Lastseen                  string        `json:"lastSeen"`
								Firstseenduration         string        `json:"firstSeenDuration"`
								Lastseenduration          string        `json:"lastSeenDuration"`
								Firstseenlastseenduration string        `json:"firstSeenLastSeenDuration"`
								Organizations             []interface{} `json:"organizations"`
								Values                    []struct {
									Value string `json:"value"`
								} `json:"values"`
							} `json:"records"`
							Previews []struct {
								Firstseen                 string        `json:"firstSeen"`
								Lastseen                  string        `json:"lastSeen"`
								Firstseenduration         string        `json:"firstSeenDuration"`
								Lastseenduration          string        `json:"lastSeenDuration"`
								Firstseenlastseenduration string        `json:"firstSeenLastSeenDuration"`
								Organizations             []interface{} `json:"organizations"`
								Values                    []struct {
									Value string `json:"value"`
								} `json:"values"`
							} `json:"previews"`
						} `json:"data"`
					} `json:"txt"`
				} `json:"history"`
			} `json:"data"`
			Asnrisklevel string `json:"asnRiskLevel"`
		} `json:"dnsData"`
		Domain           string `json:"domain"`
		Locationpathname string `json:"locationPathname"`
		Locationsearch   string `json:"locationSearch"`
		Page             int    `json:"page"`
		Searchvalue      string `json:"searchValue"`
		Subdomainscount  int    `json:"subdomainsCount"`
		User             struct {
			Email      string `json:"email"`
			Isverified bool   `json:"isVerified"`
			Package    string `json:"package"`
			Token      string `json:"token"`
		} `json:"user"`
	} `json:"pageProps"`
	NSsp bool `json:"__N_SSP"`
}

type AHistoryRespsonse struct {
	Pageprops struct {
		Apexdomain string `json:"apexDomain"`
		Dnsdata    struct {
			Success    bool   `json:"success"`
			Status     int    `json:"status"`
			Statustext string `json:"statusText"`
			Error      string `json:"error"`
			Data       struct {
				History struct {
					A struct {
						Data struct {
							Pages   int `json:"pages"`
							Records []struct {
								Firstseen                 string `json:"firstSeen"`
								Lastseen                  string `json:"lastSeen"`
								Firstseenduration         string `json:"firstSeenDuration"`
								Lastseenduration          string `json:"lastSeenDuration"`
								Firstseenlastseenduration string `json:"firstSeenLastSeenDuration"`
								Values                    []struct {
									IP      string `json:"ip"`
									IPCount int    `json:"ip_count"`
								} `json:"values"`
								Organizations []string `json:"organizations"`
							} `json:"records"`
						} `json:"data"`
					} `json:"a"`
				} `json:"history"`
			} `json:"data"`
			Asnrisklevel string `json:"asnRiskLevel"`
		} `json:"dnsData"`
		Dnstypes        []string `json:"dnsTypes"`
		Domain          string   `json:"domain"`
		Page            int      `json:"page"`
		Searchvalue     string   `json:"searchValue"`
		Subdomain       string   `json:"subdomain"`
		Subdomainscount int      `json:"subdomainsCount"`
		Type            string   `json:"type"`
		User            struct {
			Email      string `json:"email"`
			Isverified bool   `json:"isVerified"`
			Package    string `json:"package"`
			Token      string `json:"token"`
		} `json:"user"`
	} `json:"pageProps"`
	NSsp bool `json:"__N_SSP"`
}

type SameServerResponse struct {
	Pageprops struct {
		IP               string `json:"ip"`
		Islist           bool   `json:"isList"`
		Locationsearch   string `json:"locationSearch"`
		Locationpathname string `json:"locationPathname"`
		Page             int    `json:"page"`
		Searchvalue      string `json:"searchValue"`
		Serverresponse   struct {
			Success    bool   `json:"success"`
			Status     int    `json:"status"`
			Statustext string `json:"statusText"`
			Error      string `json:"error"`
			Data       struct {
				Meta struct {
					LimitReached bool `json:"limit_reached"`
					MaxPage      int  `json:"max_page"`
					Page         int  `json:"page"`
					TotalPages   int  `json:"total_pages"`
				} `json:"meta"`
				Records []struct {
					HostProvider []string      `json:"host_provider"`
					Hostname     string        `json:"hostname"`
					MailProvider []interface{} `json:"mail_provider"`
					OpenPageRank interface{}   `json:"open_page_rank"`
				} `json:"records"`
				Total    int `json:"total"`
				Previews []struct {
					Hostname string `json:"hostname"`
					Rank     string `json:"rank"`
					Provider string `json:"provider"`
				} `json:"previews"`
			} `json:"data"`
			Asnrisklevel string `json:"asnRiskLevel"`
		} `json:"serverResponse"`
		Type string `json:"type"`
		User struct {
			Email      string `json:"email"`
			Isverified bool   `json:"isVerified"`
			Package    string `json:"package"`
			Token      string `json:"token"`
		} `json:"user"`
	} `json:"pageProps"`
	NSsp bool `json:"__N_SSP"`
}
