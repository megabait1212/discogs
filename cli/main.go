// http://www.discogs.com/developers
package main

import (
	"fmt"

	"github.com/dmikalova/discogs"
)

func main() {

	client := discogs.NewClient(nil)

	// // Identity
	// identity := new(discogs.Identity)
	// identity, _, err := client.Identity.Get(311)
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// }
	// fmt.Println("Identity:	", identity.Username)

	// // User
	// user := new(discogs.User)
	// user, _, err = client.User.Get("teratomata")
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// }
	// fmt.Println("Rank:		", user.Name)

	// // User Post
	// userPost := discogs.UserPost{
	// 	Username: "teratomata",
	// 	Name:     "David",
	// }

	// user, _, err := client.User.Post("teratomata", userPost)
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// }
	// fmt.Println("Name:		", user.Name)

	// // User Submissions
	// userSubmissions := new(discogs.UserSubmissions)
	// userSubmissions, _, err = client.UserSubmissions.Get("teratomata")
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// }
	// fmt.Println("Submissions:	", userSubmissions.Pagination.Items)

	// // User Contributions
	// userContributions := new(discogs.UserContributions)
	// userContributions, _, err = client.UserContributions.Get("teratomata")
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// }
	// fmt.Println("Contributions:	", userContributions.Pagination.Items)

	// // User Inventory
	// userInventory := new(discogs.UserInventory)
	// userInventory, _, err = client.UserInventory.Get("teratomata")
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// }
	// fmt.Println("Inventory:	", userInventory.Pagination.Items)

	// // User Collection Folders
	// userCollectionFolders := new(discogs.UserCollectionFolders)
	// userCollectionFolders, _, err = client.UserCollectionFolders.Get("teratomata")
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// }
	// fmt.Println("Folders:	", userCollectionFolders.Folders[0].Name)

	// // User Collection Folder
	// userCollectionFolder := new(discogs.UserCollectionFolder)
	// userCollectionFolder, _, err = client.UserCollectionFolder.Get("teratomata", 0)
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// }
	// fmt.Println("Folder:		", userCollectionFolder.Name)

	// // User Collection Folder Releases
	// userCollectionFolderReleases := new(discogs.UserCollectionFolderReleases)
	// userCollectionFolderReleases, _, err = client.UserCollectionFolderReleases.Get("teratomata", 0)
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// }
	// fmt.Println("Folder Items:	", userCollectionFolderReleases.Pagination.Items)

	// // User Collection Folder Releases
	// userCollectionFields := new(discogs.UserCollectionFields)
	// userCollectionFields, _, err = client.UserCollectionFields.Get("teratomata")
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// }
	// fmt.Println("Fields:		", userCollectionFields.Fields[0].Name)

	// // User Wantlist
	// userWantList := new(discogs.UserWants)
	// userWantList, _, err = client.UserWants.Get("teratomata")
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// }
	// fmt.Println("Wants:		", userWantList.Pagination.Items)

	// // Marketplace Listing
	// marketplaceListing := new(discogs.MarketplaceListing)
	// marketplaceListing, _, err = client.MarketplaceListing.Get(120063622)
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// }
	// fmt.Println("Listing:	", marketplaceListing.Price.Value)

	// Marketplace Listing New
	marketplaceListingEdit := discogs.MarketplaceListingEdit{
		ReleaseID: 208286,
		Condition: "Fair (F)",
		Price:     60,
		Status:    "Draft",
	}

	newListing, _, err := client.MarketplaceListing.New(marketplaceListingEdit)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("New Listing:	", newListing.ListingID)

	// // Marketplace Listing Edit
	// marketplaceListingEdit := discogs.MarketplaceListingEdit{
	// 	ReleaseID: 208286,
	// 	Condition: "Fair (F)",
	// 	Price:     60,
	// 	Status:    "Draft",
	// }

	// _, err = client.MarketplaceListing.Edit(202293354, marketplaceListingEdit)
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// }

	// // Marketplace Listing Delete
	// marketplaceListingDelete := discogs.MarketplaceListingDelete{
	// 	ListingID: 202293354,
	// }

	// _, err := client.MarketplaceListing.Delete(202293354, marketplaceListingDelete)
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// }

	fmt.Println(discogs.Condition.Mint())
	fmt.Println(discogs.SleeveCondition.Poor())

	// // Marketplace Order
	// // marketplaceOrder := new(discogs.MarketplaceOrder)
	// // marketplaceOrder, _, err = client.MarketplaceOrder.Get(0)
	// // if err != nil {
	// // 	fmt.Println("Error: ", err)
	// // }
	// // fmt.Println("Order:	", marketplaceOrder.Status)

	// // Marketplace Order Messages
	// // marketplaceOrderMessages := new(discogs.MarketplaceOrderMessages)
	// // marketplaceOrderMessages, _, err = client.MarketplaceOrderMessages.Get(0)
	// // if err != nil {
	// // 	fmt.Println("Error: ", err)
	// // }
	// // fmt.Println("Messages:	", marketplaceOrderMessages.Pagination.Items)

	// // Marketplace Fee
	// marketplaceFee := new(discogs.MarketplaceFee)
	// marketplaceFee, _, err = client.MarketplaceFee.Get(10, "USD")
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// }
	// fmt.Println("Fee:		", "$", marketplaceFee.Value, marketplaceFee.Currency)

	// // Marketplace Price Suggestions
	// // marketplacePriceSuggestions := new(discogs.MarketplacePriceSuggestions)
	// // marketplacePriceSuggestions, _, err = client.MarketplacePriceSuggestions.Get(208286)
	// // if err != nil {
	// // 	fmt.Println("Error: ", err)
	// // }
	// // fmt.Println("Price:	$", marketplacePriceSuggestions.VeryGood.Value, marketplacePriceSuggestions.VeryGood.Currency)

	// // Label
	// label := new(discogs.Label)
	// label, _, err = client.Label.Get(311)
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// }
	// fmt.Println("Label:		", label.Name)

	// // Label Releases
	// labelReleases := new(discogs.LabelReleases)
	// labelReleases, _, err = client.LabelReleases.Get(311)
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// }
	// fmt.Println("Releases:	", labelReleases.Pagination.Items)

	// // Artist
	// artist := new(discogs.Artist)
	// artist, _, err = client.Artist.Get(1373)
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// }
	// fmt.Println("Artist:		", artist.Name)

	// // Artist Releases
	// artistReleases := new(discogs.ArtistReleases)
	// artistReleases, _, err = client.ArtistReleases.Get(1373)
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// }
	// fmt.Println("Releases:	", artistReleases.Pagination.Items)

	// // Master
	// master := new(discogs.Master)
	// master, _, err = client.Master.Get(2815)
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// }
	// fmt.Println("Release:	", master.Title)

	// // Cover
	// image := new(discogs.Image)
	// image, _, err = client.Image.Get("http://api.discogs.com/images/R-208286-1162232459.jpeg")
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// }
	// fmt.Println("Cover:		", image)

	// // Master Versions
	// masterVersions := new(discogs.MasterVersions)
	// masterVersions, _, err = client.MasterVersions.Get(2815)
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// }
	// fmt.Println("Versions:	", masterVersions.Pagination.Items)

	// Release
	release := new(discogs.Release)
	release, _, err = client.Release.Get(208286)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("Version:	", release.Title)
	fmt.Println("Tracks:		", len(release.Tracklist))
}
