package rappelconso

import "github.com/xefiji/rappelconso/pkg/datagouv"

type RecordEntity struct {
	ID             string `db:"id"`
	UUID           string `db:"uuid"`
	Timestamp      string `db:"timestamp"`
	Size           int    `db:"size"`
	SelfLink       string `db:"self_link"`
	Ref            string `db:"reference_fiche"`
	Nature         string `db:"nature_juridique_du_rappel"`
	Cat            string `db:"categorie_de_produit"`
	SubCategory    string `db:"sous_categorie_de_produit"`
	Brand          string `db:"nom_de_la_marque_du_produit"`
	Model          string `db:"noms_des_modeles_ou_references"`
	ProductID      string `db:"identification_des_produits"`
	Packaging      string `db:"conditionnements"`
	SellStartEnd   string `db:"date_debut_fin_de_commercialisation"`
	Infos          string `db:"informations_complementaires"`
	GeoLoc         string `db:"zone_geographique_de_vente"`
	Seller         string `db:"distributeurs"`
	Reason         string `db:"motif_du_rappel"`
	Risks          string `db:"risques_encourus_par_le_consommateur"`
	SanitaryAdvice string `db:"preconisations_sanitaires"`
	ConsoAdvice    string `db:"conduites_a_tenir_par_le_consommateur"`
	Contact        string `db:"numero_de_contact"`
	Compensation   string `db:"modalites_de_compensation"`
	RecallEnd      string `db:"date_de_fin_de_la_procedure_de_rappel"`
	PublicInfo     string `db:"informations_complementaires_publiques"`
	Images         string `db:"lien_vers_la_liste_des_images"`
	PDF            string `db:"lien_vers_l_affichette_pdf"`
	DateRef        string `db:"date_ref"`
}

func newRecordEntity(r datagouv.RecordData) RecordEntity {
	var selfLink string
	for _, link := range r.Links {
		if link.Rel == "self" {
			selfLink = link.Href
		}
	}
	re := RecordEntity{
		UUID:           r.Record.ID,
		Timestamp:      r.Record.Timestamp,
		Size:           r.Record.Size,
		SelfLink:       selfLink,
		Ref:            r.Record.Fields.Ref,
		Nature:         r.Record.Fields.Nature,
		Cat:            r.Record.Fields.Cat,
		SubCategory:    r.Record.Fields.SubCategory,
		Brand:          r.Record.Fields.Brand,
		Model:          r.Record.Fields.Model,
		ProductID:      r.Record.Fields.ProductID,
		Packaging:      r.Record.Fields.Packaging,
		SellStartEnd:   r.Record.Fields.SellStartEnd,
		Infos:          r.Record.Fields.Infos,
		GeoLoc:         r.Record.Fields.GeoLoc,
		Seller:         r.Record.Fields.Seller,
		Reason:         r.Record.Fields.Reason,
		Risks:          r.Record.Fields.Risks,
		SanitaryAdvice: r.Record.Fields.SanitaryAdvice,
		ConsoAdvice:    r.Record.Fields.ConsoAdvice,
		Contact:        r.Record.Fields.Contact,
		Compensation:   r.Record.Fields.Compensation,
		RecallEnd:      r.Record.Fields.RecallEnd,
		PublicInfo:     r.Record.Fields.PublicInfo,
		Images:         r.Record.Fields.Images,
		PDF:            r.Record.Fields.PDF,
		DateRef:        r.Record.Fields.DateRef,
	}
	return re
}
