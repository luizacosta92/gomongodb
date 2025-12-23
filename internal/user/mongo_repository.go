package user

//O que é: implementação do UserRepository com o driver oficial.
//Por que importa: concentra a lógica de acesso ao Mongo; o resto do sistema não sabe “como” os dados são buscados.
import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	mongoDriver "go.mongodb.org/mongo-driver/mongo"
)

// Declara o que deve ter no repositório
type MongoUserRepository struct {
	collection *mongoDriver.Collection
}

// Inicialização do repositório
func NewMongoRepository(collection *mongoDriver.Collection) *MongoUserRepository {
	return &MongoUserRepository{collection: collection}
}

//criar funcoes para o repositório:
//create
//findall
//findbydpp
//findbysite
//findbyage

//Uma coleção pode ser usada para consultar o banco de dados ou inserir documentos:

func (r *MongoUserRepository) Create(ctx context.Context, user User) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	result, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return "", err
	}

	return idAsString(result.InsertedID), nil
}

func idAsString(id any) string {
	switch v := id.(type) {
	case string:
		return v
	default:
		return fmt.Sprintf("%v", v)
	}
}

func (r *MongoUserRepository) FindAll(ctx context.Context) ([]User, error) {
	
	//cursor para iterar sobre os resultados
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []User
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	
	return users, nil
}
//TO DO: falta implementar as outras funções: findbydpp, findbysite, findbyage

func (r *MongoUserRepository) FindByDpp(ctx context.Context, dpp time.Time) ([]User, error) {
	dayStart := time.Date(dpp.Year(), dpp.Month(), dpp.Day(), 0, 0, 0, 0, time.UTC)
	dayEnd := dayStart.Add(24 * time.Hour)
	filter := bson.M{"dpp": bson.M{"$gte": dayStart, "$lte": dayEnd}}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []User
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *MongoUserRepository) FindByCity(ctx context.Context, city string) ([]User, error) {
 
	filter := bson.M{"city": city}
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []User
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *MongoUserRepository) FindByAge(ctx context.Context, age int) ([]User, error) {
	 

	filter := bson.M{"age": age}
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []User
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}