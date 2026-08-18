package storage

import "api-products/models"

var Products = []models.Product{
	{ID: 1, Name: "Monitor", Price: 700},
	{ID: 2, Name: "Teclado", Price: 150},
}

var NextID = 3
