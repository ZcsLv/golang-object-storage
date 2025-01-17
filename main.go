package main

import (
	"fmt"
	"time"

	"github.com/olivere/elastic/v7"
)

type Tweet struct {
	User     string                `json:"user"`
	Message  string                `json:"message"`
	Retweets int                   `json:"retweets"`
	Image    string                `json:"image,omitempty"`
	Created  time.Time             `json:"created,omitempty"`
	Tags     []string              `json:"tags,omitempty"`
	Location string                `json:"location,omitempty"`
	Suggest  *elastic.SuggestField `json:"suggest_field,omitempty"`
}

func main() {
	//errorlog := log.New(os.Stdout, "APP ", log.LstdFlags)

	// Obtain a client. You can also provide your own HTTP client here.
	//urlOpt := elastic.SetURL("http://localhost:9200")
	//client, err := elastic.NewClient(urlOpt)
	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetSniff(false),
	)
	fmt.Println(client, err)
	// Trace request and response details like this
	// client, err := elastic.NewClient(elastic.SetTraceLog(log.New(os.Stdout, "", 0)))
	//	if err != nil {
	//		// Handle error
	//		panic(err)
	//	}
	//
	//	// Ping the Elasticsearch server to get e.g. the version number
	//	info, code, err := client.Ping("http://127.0.0.1:9200").Do(context.Background())
	//	if err != nil {
	//		// Handle error
	//		panic(err)
	//	}
	//	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
	//
	//	// Getting the ES version number is quite common, so there's a shortcut
	//	esversion, err := client.ElasticsearchVersion("http://127.0.0.1:9200")
	//	if err != nil {
	//		// Handle error
	//		panic(err)
	//	}
	//	fmt.Printf("Elasticsearch version %s\n", esversion)
	//
	//	// Use the IndexExists service to check if a specified index exists.
	//	exists, err := client.IndexExists("twitter").Do(context.Background())
	//	if err != nil {
	//		// Handle error
	//		panic(err)
	//	}
	//	if !exists {
	//		// Create a new index.
	//		mapping := `
	//{
	//	"settings":{
	//		"number_of_shards":1,
	//		"number_of_replicas":0
	//	},
	//	"mappings":{
	//		"doc":{
	//			"properties":{
	//				"user":{
	//					"type":"keyword"
	//				},
	//				"message":{
	//					"type":"text",
	//					"store": true,
	//					"fielddata": true
	//				},
	//                "retweets":{
	//                    "type":"long"
	//                },
	//				"tags":{
	//					"type":"keyword"
	//				},
	//				"location":{
	//					"type":"geo_point"
	//				},
	//				"suggest_field":{
	//					"type":"completion"
	//				}
	//			}
	//		}
	//	}
	//}
	//`
	//		createIndex, err := client.CreateIndex("twitter").Body(mapping).IncludeTypeName(true).Do(context.Background())
	//		if err != nil {
	//			// Handle error
	//			panic(err)
	//		}
	//		if !createIndex.Acknowledged {
	//			// Not acknowledged
	//		}
	//	}
	//
	//	// Index a tweet (using JSON serialization)
	//	tweet1 := Tweet{User: "olivere", Message: "Take Five", Retweets: 0}
	//	put1, err := client.Index().
	//		Index("twitter").
	//		Type("doc").
	//		Id("1").
	//		BodyJson(tweet1).
	//		Do(context.Background())
	//	if err != nil {
	//		// Handle error
	//		panic(err)
	//	}
	//	fmt.Printf("Indexed tweet %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
	//
	//	// Index a second tweet (by string)
	//	tweet2 := `{"user" : "olivere", "message" : "It's a Raggy Waltz"}`
	//	put2, err := client.Index().
	//		Index("twitter").
	//		Type("doc").
	//		Id("2").
	//		BodyString(tweet2).
	//		Do(context.Background())
	//	if err != nil {
	//		// Handle error
	//		panic(err)
	//	}
	//	fmt.Printf("Indexed tweet %s to index %s, type %s\n", put2.Id, put2.Index, put2.Type)
	//
	//	// Get tweet with specified ID
	//	get1, err := client.Get().
	//		Index("twitter").
	//		Type("doc").
	//		Id("1").
	//		Do(context.Background())
	//	if err != nil {
	//		switch {
	//		case elastic.IsNotFound(err):
	//			panic(fmt.Sprintf("Document not found: %v", err))
	//		case elastic.IsTimeout(err):
	//			panic(fmt.Sprintf("Timeout retrieving document: %v", err))
	//		case elastic.IsConnErr(err):
	//			panic(fmt.Sprintf("Connection problem: %v", err))
	//		default:
	//			// Some other kind of error
	//			panic(err)
	//		}
	//	}
	//	fmt.Printf("Got document %s in version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)
	//
	//	// Flush to make sure the documents got written.
	//	_, err = client.Flush().Index("twitter").Do(context.Background())
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	// Search with a term query
	//	termQuery := elastic.NewTermQuery("user", "olivere")
	//	searchResult, err := client.Search().
	//		Index("twitter").        // search in index "twitter"
	//		Query(termQuery).        // specify the query
	//		Sort("user", true).      // sort by "user" field, ascending
	//		From(0).Size(10).        // take documents 0-9
	//		Pretty(true).            // pretty print request and response JSON
	//		Do(context.Background()) // execute
	//	if err != nil {
	//		// Handle error
	//		panic(err)
	//	}
	//
	//	// searchResult is of type SearchResult and returns hits, suggestions,
	//	// and all kinds of other information from Elasticsearch.
	//	fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)
	//
	//	// Each is a convenience function that iterates over hits in a search result.
	//	// It makes sure you don't need to check for nil values in the response.
	//	// However, it ignores errors in serialization. If you want full control
	//	// over iterating the hits, see below.
	//	var ttyp Tweet
	//	for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
	//		t := item.(Tweet)
	//		fmt.Printf("Tweet by %s: %s\n", t.User, t.Message)
	//	}
	//	// TotalHits is another convenience function that works even when something goes wrong.
	//	fmt.Printf("Found a total of %d tweets\n", searchResult.TotalHits())
	//
	//	// Here's how you iterate through results with full control over each step.
	//	if searchResult.Hits.TotalHits > 0 {
	//		fmt.Printf("Found a total of %d tweets\n", searchResult.Hits.TotalHits)
	//
	//		// Iterate through results
	//		for _, hit := range searchResult.Hits.Hits {
	//			// hit.Index contains the name of the index
	//
	//			// Deserialize hit.Source into a Tweet (could also be just a map[string]interface{}).
	//			var t Tweet
	//			err := json.Unmarshal(*hit.Source, &t)
	//			if err != nil {
	//				// Deserialization failed
	//			}
	//
	//			// Work with tweet
	//			fmt.Printf("Tweet by %s: %s\n", t.User, t.Message)
	//		}
	//	} else {
	//		// No hits
	//		fmt.Print("Found no tweets\n")
	//	}
	//
	//	// Update a tweet by the update API of Elasticsearch.
	//	// We just increment the number of retweets.
	//	script := elastic.NewScript("ctx._source.retweets += params.num").Param("num", 1)
	//	update, err := client.Update().Index("twitter").Type("doc").Id("1").
	//		Script(script).
	//		Upsert(map[string]interface{}{"retweets": 0}).
	//		Do(context.Background())
	//	if err != nil {
	//		// Handle error
	//		panic(err)
	//	}
	//	fmt.Printf("New version of tweet %q is now %d", update.Id, update.Version)
	//
	//	// ...
	//
	//	// Delete an index.
	//	deleteIndex, err := client.DeleteIndex("twitter").Do(context.Background())
	//	if err != nil {
	//		// Handle error
	//		panic(err)
	//	}
	//	if !deleteIndex.Acknowledged {
	//		// Not acknowledged
	//	}
}
