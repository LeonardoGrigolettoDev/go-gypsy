package repositories

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GenericRepository é um repositório genérico para qualquer coleção
type GenericRepository struct {
	collection *mongo.Collection
}

// NewGenericRepository cria um novo repositório genérico para a coleção especificada
func NewGenericRepository(client *mongo.Client, dbName, collectionName string) *GenericRepository {
	// Cria uma coleção no banco de dados
	collection := client.Database(dbName).Collection(collectionName)
	return &GenericRepository{collection: collection}
}

// Create insere um novo documento na coleção
func (repo *GenericRepository) Create(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error) {
	result, err := repo.collection.InsertOne(ctx, document)
	if err != nil {
		return nil, fmt.Errorf("erro ao inserir o documento: %v", err)
	}
	return result, nil
}

// GetByID busca um documento pela chave ID
func (repo *GenericRepository) GetByID(ctx context.Context, id primitive.ObjectID, result interface{}) error {
	err := repo.collection.FindOne(ctx, bson.M{"_id": id}).Decode(result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil // Nenhum documento encontrado
		}
		return fmt.Errorf("erro ao buscar o documento: %v", err)
	}
	return nil
}

// Find busca documentos que atendem a um filtro
func (repo *GenericRepository) Find(ctx context.Context, filter interface{}, result interface{}) error {
	cursor, err := repo.collection.Find(ctx, filter)
	if err != nil {
		return fmt.Errorf("erro ao buscar documentos: %v", err)
	}
	defer cursor.Close(ctx)

	// Decodifica os documentos encontrados
	for cursor.Next(ctx) {
		var elem interface{}
		if err := cursor.Decode(&elem); err != nil {
			return err
		}
		result = append(result.([]interface{}), elem)
	}

	if err := cursor.Err(); err != nil {
		return err
	}

	return nil
}

// Update atualiza um documento na coleção
func (repo *GenericRepository) Update(ctx context.Context, id primitive.ObjectID, update interface{}) (*mongo.UpdateResult, error) {
	result, err := repo.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.D{
		{"$set", update},
	})
	if err != nil {
		return nil, fmt.Errorf("erro ao atualizar o documento: %v", err)
	}
	return result, nil
}

// Delete remove um documento da coleção
func (repo *GenericRepository) Delete(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error) {
	result, err := repo.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return nil, fmt.Errorf("erro ao deletar o documento: %v", err)
	}
	return result, nil
}
