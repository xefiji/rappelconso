package rappelconso

import (
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func newRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) save(record RecordEntity) error {
	const query = `
	INSERT INTO record (
		uuid,
		timestamp,
		size,
		self_link,
		reference_fiche,
		nature_juridique_du_rappel,
		categorie_de_produit,
		sous_categorie_de_produit,
		nom_de_la_marque_du_produit,
		noms_des_modeles_ou_references,
		identification_des_produits,
		conditionnements,
		date_debut_fin_de_commercialisation,
		informations_complementaires,
		zone_geographique_de_vente,
		distributeurs,
		motif_du_rappel,
		risques_encourus_par_le_consommateur,
		preconisations_sanitaires,
		conduites_a_tenir_par_le_consommateur,
		numero_de_contact,
		modalites_de_compensation,
		date_de_fin_de_la_procedure_de_rappel,
		informations_complementaires_publiques,
		lien_vers_la_liste_des_images,
		lien_vers_l_affichette_pdf,
		date_ref
	)
	
	VALUES(
		:uuid,
		:timestamp,
		:size,
		:self_link,
		:reference_fiche,
		:nature_juridique_du_rappel,
		:categorie_de_produit,
		:sous_categorie_de_produit,
		:nom_de_la_marque_du_produit,
		:noms_des_modeles_ou_references,
		:identification_des_produits,
		:conditionnements,
		:date_debut_fin_de_commercialisation,
		:informations_complementaires,
		:zone_geographique_de_vente,
		:distributeurs,
		:motif_du_rappel,
		:risques_encourus_par_le_consommateur,
		:preconisations_sanitaires,
		:conduites_a_tenir_par_le_consommateur,
		:numero_de_contact,
		:modalites_de_compensation,
		:date_de_fin_de_la_procedure_de_rappel,
		:informations_complementaires_publiques,
		:lien_vers_la_liste_des_images,
		:lien_vers_l_affichette_pdf,
		:date_ref
	) ON CONFLICT DO NOTHING`

	rows, err := r.db.NamedQuery(query, record)
	if err != nil {
		return err
	}

	if rows != nil {
		defer rows.Close()
	}

	return nil
}

// func (r *Repository) getLasts(record RecordEntity) ([]*RecordEntity, error) {
// 	return nil, nil
// }

func (r *Repository) countDBRecords() (int, error) {
	count := 0
	const query = `SELECT count(id) FROM record`

	return count, r.db.Get(&count, query)
}

func (r *Repository) flushDBRecords() error {
	if _, err := r.db.Exec(`TRUNCATE TABLE record`); err != nil {
		return err
	}

	if _, err := r.db.Exec(`ALTER SEQUENCE record_id_seq RESTART WITH 1`); err != nil {
		return err
	}

	return nil
}
