CREATE TABLE IF NOT EXISTS "record" (
    "id" SERIAL PRIMARY KEY,
    "uuid" character varying(255) NOT NULL UNIQUE,
    "timestamp" timestamp NOT NULL,
    "size" integer NOT NULL,
    "reference_fiche" character varying(255),
    "alert_twitter" timestamp,
    "alert_facebook" timestamp,
    "nature_juridique_du_rappel" TEXT,
    "categorie_de_produit" TEXT,
    "sous_categorie_de_produit" TEXT,
    "nom_de_la_marque_du_produit" TEXT,
    "noms_des_modeles_ou_references" TEXT,
    "identification_des_produits" TEXT,
    "conditionnements" TEXT,
    "date_debut_fin_de_commercialisation" TEXT,
    "informations_complementaires" TEXT,
    "zone_geographique_de_vente" TEXT,
    "distributeurs" TEXT,
    "motif_du_rappel" TEXT,
    "risques_encourus_par_le_consommateur" TEXT,
    "preconisations_sanitaires" TEXT,
    "conduites_a_tenir_par_le_consommateur" TEXT,
    "numero_de_contact" TEXT,
    "modalites_de_compensation" TEXT,
    "date_de_fin_de_la_procedure_de_rappel" TEXT,
    "informations_complementaires_publiques" TEXT,
    "self_link" TEXT,
    "lien_vers_la_liste_des_images" TEXT,
    "lien_vers_l_affichette_pdf" TEXT,
    "date_ref"  character varying(255)
) WITH (oids = false);

CREATE INDEX ON record (uuid);

---- create above / drop below ----

DROP TABLE IF EXISTS "record";
