package datagouv

type Records struct {
	Total      uint         `json:"total_count"`
	Links      []Link       `json:"links"`
	RecordData []RecordData `json:"records"`
}

type Link struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}

type RecordData struct {
	Links  []Link `json:"links"`
	Record Record `json:"record"`
}

type Record struct {
	ID        string `json:"id"`
	Timestamp string `json:"timestamp"`
	Size      int    `json:"size"`
	Fields    Fields `json:"fields"`
}

type Fields struct {
	Ref            string `json:"reference_fiche"`
	Nature         string `json:"nature_juridique_du_rappel"`
	Cat            string `json:"categorie_de_produit"`
	SubCategory    string `json:"sous_categorie_de_produit"`
	Brand          string `json:"nom_de_la_marque_du_produit"`
	Model          string `json:"noms_des_modeles_ou_references"`
	ProductID      string `json:"identification_des_produits"`
	Packaging      string `json:"conditionnements"`
	SellStartEnd   string `json:"date_debut_fin_de_commercialisation"`
	Infos          string `json:"informations_complementaires"`
	GeoLoc         string `json:"zone_geographique_de_vente"`
	Seller         string `json:"distributeurs"`
	Reason         string `json:"motif_du_rappel"`
	Risks          string `json:"risques_encourus_par_le_consommateur"`
	SanitaryAdvice string `json:"preconisations_sanitaires"`
	ConsoAdvice    string `json:"conduites_a_tenir_par_le_consommateur"`
	Contact        string `json:"numero_de_contact"`
	Compensation   string `json:"modalites_de_compensation"`
	RecallEnd      string `json:"date_de_fin_de_la_procedure_de_rappel"`
	PublicInfo     string `json:"informations_complementaires_publiques"`
	Images         string `json:"lien_vers_la_liste_des_images"`
	PDF            string `json:"lien_vers_l_affichette_pdf"`
	DateRef        string `json:"date_ref"`
}
