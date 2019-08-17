package main

import(
  "encoding/json"
  "log"
  "net/http"
  "math/rand"
  "strconv"
  "github.com/gorilla/mux"
)

// Item Struct
type Item struct {
  ID      string `json:"id"`
  Item    string `json:"item"`
  Price   string `json:"price"`
  Age     string `json:"age"`
}

// Init Items Variable as Slice Item Struct
var items []Item

// Get All Items
func getItems(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(items)
}

// Get Single Item
func getItem(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r) // Get params

  // Iterate Over Items, Find ID
  for _, item := range items {
    if item.ID == params["id"] {
      json.NewEncoder(w).Encode(item)
      return
    }
  }
  json.NewEncoder(w).Encode(&Item{})
}

// Create a new Item
func createItem(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")
  var item Item
  _ = json.NewDecoder(r.Body).Decode(&item)
  item.ID = strconv.Itoa(rand.Intn(10000000)) // Mock ID -- Not for production
  items = append(items, item)
  json.NewEncoder(w).Encode(&item)
}

// Update Item
func updateItem(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  for index, item := range items {
    if item.ID == params["id"]{
      items = append(items[:index], items[index+1:]...)
      var item Item
      _ = json.NewDecoder(r.Body).Decode(&item)
      item.ID = params["id"]
      items = append(items, item)
      json.NewEncoder(w).Encode(&item)
      return
    }
  }
  json.NewEncoder(w).Encode(items)
}

// Delete Item
func deleteItem(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  for index, item := range items {
    if item.ID == params["id"]{
      items = append(items[:index], items[index+1:]...)
      break
    }
  }
  json.NewEncoder(w).Encode(items)
}

func main() {
  // Initialize Router
  r := mux.NewRouter()

  // Mock Data - TODO: implement DB
  items = append(items, Item{ID: "1", Item: "Computer", Price: "1525.00", Age: "5"}) // Age in years
  items = append(items, Item{ID: "2", Item: "Telescope", Price: "2250.00", Age: "2"})
  items = append(items, Item{ID: "3", Item: "Guitar", Price: "125.00", Age: "7"})
  
  // Route Handlers / API Endpoints
  r.HandleFunc("/api/items", getItems).Methods("GET")
  r.HandleFunc("/api/items/{id}", getItem).Methods("GET")
  r.HandleFunc("/api/items", createItem).Methods("POST")
  r.HandleFunc("/api/items/{id}", updateItem).Methods("PUT")
  r.HandleFunc("/api/items/{id}", deleteItem).Methods("DELETE")

  log.Fatal(http.ListenAndServe(":8000", r)) // Start Server & throw error if fatal
}
