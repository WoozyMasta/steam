/*
Package appid provides a collection of constants representing Steam application IDs.

These constants encompass a variety of games across different engines and categories,
including those built on the GoldSource and Source 1/2 engines, as well as other popular titles.

References:
  - https://api.steampowered.com/ISteamApps/GetAppList/v2/
  - https://api.steampowered.com/IStoreService/GetAppList/v1/
  - https://steamdb.info/
*/
package appid

import "fmt"

//go:generate go run generate.go

// AppID represents a Steam application ID.
type AppID uint64

const (
	Unknown AppID = 0 // Any unexpected app

	// GoldSource
	CounterStrike       AppID = 10  // Counter Strike 1.6
	TeamFortressClassic AppID = 20  // Team Fortress Classic
	DayOfDefeat         AppID = 30  // Day Of Defeat
	DeathMatchClassic   AppID = 40  // Half-Life: DeathMatch Classic
	OpposingForce       AppID = 50  // Opposing Force
	Ricochet            AppID = 60  // Ricochet
	HalfLife            AppID = 70  // Half-Life
	CounterStrikeCZ     AppID = 80  // Counter Strike: Condition Zero
	HalfLifeBlueShift   AppID = 130 // Half-Life: BlueShift

	// Source 1/2

	HalfLife2           AppID = 220     // Half-Life 2
	CounterStrikeSource AppID = 240     // Counter Strike Source
	DayOfDefeatSource   AppID = 300     // Day Of Defeat Source
	HalfLifeDMSource    AppID = 360     // Half-Life: DeathMatch Source
	Portal              AppID = 400     // Portal
	TeamFortress2       AppID = 440     // Team Fortress 2
	Spacewar            AppID = 480     // Spacewar
	Left4Dead           AppID = 500     // Left 4 Dead
	Left4Dead2          AppID = 550     // Left 4 Dead 2
	Dota2               AppID = 570     // Dota 2
	Portal2             AppID = 620     // Portal 2
	AlienSwarm          AppID = 630     // Alien Swarm
	CounterStrike2      AppID = 730     // Counter Strike 2
	KillingFloor        AppID = 1250    // Killing Floor
	TheShip             AppID = 2400    // The Ship
	GarrysMod           AppID = 4000    // Garry's Mod
	Deadlock            AppID = 1422450 // Deadlock

	// Other

	Arma3                    AppID = 107410  // Arma 3
	DayZ                     AppID = 221100  // DayZ
	DayZExp                  AppID = 1024020 // DayZ Experimental
	Rust                     AppID = 252490  // Rust
	MarvelRivals             AppID = 2767030 // Marvel Rivals
	PathOfExile2             AppID = 2694490 // Path of Exile 2
	PubgBattlegrounds        AppID = 578080  // PUBG: BATTLEGROUNDS
	GTA5                     AppID = 271590  // Grand Theft Auto V
	Palworld                 AppID = 1623730 // Palworld
	BaldusGate3              AppID = 1086940 // Baldur's Gate 3
	CallOfDuty               AppID = 1938090 // Call of Duty
	StardewValley            AppID = 413150  // Stardew Valley
	RainbowSixSiege          AppID = 359550  // Tom Clancy's Rainbow Six Siege
	DeltaForce               AppID = 2507950 // Delta Force
	HellDivers2              AppID = 553850  // HELLDIVERS™ 2
	Warframe                 AppID = 230410  // Warframe
	EldenRing                AppID = 1245620 // ELDEN RING
	ApexLegends              AppID = 1172470 // Apex Legends
	WarThunder               AppID = 236390  // War Thunder
	CivilizationVI           AppID = 289070  // Sid Meier's Civilization VI
	Cyberpunk2077            AppID = 1091500 // Cyberpunk 2077
	VRChat                   AppID = 438100  // VRChat
	FootballManager2024      AppID = 2252570 // Football Manager 2024
	FC25                     AppID = 2669320 // EA SPORTS FC 25
	ProjectZomboid           AppID = 108600  // Project Zomboid
	ThroneAndLiberty         AppID = 2429640 // THRONE AND LIBERTY
	HeartsOfIronIV           AppID = 394360  // Hearts of Iron IV
	TheSims4                 AppID = 1222670 // The Sims™ 4
	RedDeadRedemption2       AppID = 1174180 // Red Dead Redemption 2
	LostArk                  AppID = 1599340 // Lost Ark
	ARKSurvivalAscended      AppID = 2399830 // ARK: Survival Ascended
	Phasmophobia             AppID = 739630  // Phasmophobia
	SevenDaysToDie           AppID = 251570  // 7 Days to Die
	PayDay2                  AppID = 218620  // PAYDAY 2
	DeadByDaylight           AppID = 381210  // Dead by Daylight
	Factorio                 AppID = 427520  // Factorio
	Terraria                 AppID = 105600  // Terraria
	TheBindingOfIsaacRebirth AppID = 250900  // The Binding of Isaac: Rebirth
	TotalWarWarhammerIII     AppID = 1142710 // Total War: WARHAMMER III
	ARKSurvivalEvolved       AppID = 346110  // ARK: Survival Evolved
	RimWorld                 AppID = 294100  // RimWorld
	RocketLeague             AppID = 252950  // Rocket League
	Satisfactory             AppID = 526870  // Satisfactory
	EuroTruckSimulator2      AppID = 227300  // Euro Truck Simulator 2
	BlackDesert              AppID = 582660  // Black Desert
	Destiny2                 AppID = 1085660 // Destiny 2
	DontStarveTogether       AppID = 322330  // Don't Starve Together
	BeamNGdrive              AppID = 2841600 // BeamNG.drive
	SCUM                     AppID = 513710  // SCUM
	Squad                    AppID = 393380  // Squad
)

// String returns name if exists or ID for the given AppID.
func (id AppID) String() string {
	if name, ok := AppName[id]; ok {
		return name
	}

	return fmt.Sprint(uint64(id))
}
