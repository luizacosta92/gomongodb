package mongo

import (
	"context"

	mongoDriver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
/*O que é: utilitário para conectar/pingar o Mongo e obter uma coleção.
Por que importa: encapsula acesso a infra (timeouts, ping), reduz acoplamento e facilita trocar implementação depois.
*/
func Connect(ctx context.Context, uri string) (*mongoDriver.Client, error) {
	client, err := mongoDriver.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	return client, nil
}

func Collection(client *mongoDriver.Client, dbName string, collectionName string) *mongoDriver.Collection {
	return client.Database(dbName).Collection(collectionName)
}