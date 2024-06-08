# Haversack Roadmap

The development of Haversack will be divided into phases, each focusing on a category of features that are important to different types of users. Haversack will be in alpha until the completion of Phase 1, which will be the minimum viable product for the app (a basic character sheet and dice roller). The beta will run through phases 2 and 3, after which the app will be ready for a full release.

## Phase One: Character Sheet

A user will be able to use Haversack to create and reference one or several Dungeons and Dragons 5e characters. They'll be able to import a character, create one from scratch, or use a builder to guide them in character creation. Stats that are determined by other stats, like initiative and skill bonuses, will be calculated by the program, but any stat can be manually overridden. The program will not support guided levelling up in this phase.

### Requirements and Features

- **Dice**: A dice roller that can parse dice notation (e.g. `2d6+3`) and return a result. Should also support variations like advantage and disadvantage. Also include a flat D20 function.
- **Heros**: A structure that holds all the information about a character at the time of creation. Base stats, skills, proficiencies, etc. should all be stored here. This will link to other structures like classes.
- **Hero Status**: Specific information at a point in time about a character. Details like health, wealth, spell slots, inventory, etc. will be stored here. This will link to the hero structure. Should also be extensible to allow for class-specific data, like monks' ki points or barbarians' rages.
- **Classes**: Each of the classes in D&D 5e will be represented as a structure that holds information about the class. This will include things like hit dice, proficiencies, and class features. This will link to the hero structure. This also includes any special features that are class-specific, like channel divinity or sneak attack.
- **Spell Slots**: A structure that holds information about a character's spell slots. This will link to the hero status structure.
- **Consumable Features**: A generic structure to represent any feature that can be used in game that is expended and recharged. Each class will have a collection of these features within their hero structure, and statuses for each in their hero status structure. This should be as generic as possible. Anything that can be used and recharged, like spell slots, rage, ki, class features, etc. should be represented here. If it says something like "once per long rest," "regain 1d4 uses on a short rest", "3/day," it should be here. Add methods to reload and consume these features. One of the structure fields should be a callback function that is called when the feature is consumed, and another should be a check function that runs before the item is consumed to see if it's possible or allowed.
- **Game Time**: A structure for tracking the passage of time in the game. Durations will be stored in a format that can combine any number of any of the time increments. Any durations longer than a day will be handled by a separate calendar system that will be introduced in Phase Two. Bonus actions and reactions are considered instantaneous. The units of time are as follows:
  - Second: (base unit)
  - Action: 6 seconds
  - Round: 6 seconds
  - Minute: 10 rounds
  - Hour: 60 minutes
  - Day: 24 hours
- **Game Actions**: A system for handling game actions. Eventually, this will heavily prioritize combat actions, but should be extensible to handle any action that can be taken in the game. This will include things like attacks, spells, and skill checks. Game actions have one or more targets, a source, a duration, and one or more results. Results can modify any parameter of a hero or creature's status. For example, if your hero, Mitch, casts eldritch blast on a spider, here's an example of what that game action might look like:
```jsonc
// pseudocode
gameAction = {
    "source": "Mitch",
    "targets": ["Spider"],
    "duration": { "action": 1 },
    "results": [
        damageTarget({
            "target": "Spider",
            "damage": 10,
            "type": "force",
            "source": "Mitch"
        })
        // other results
    ]
}
```
- **Save Progress**: Your hero, status, and all other data can be saved locally. Haversack will be entirely local, except for retreiving new game reference materials (phase two) and networking with the rest of your party (phase five). Data will be saved in a local SQLite database.

## Phase Two: Reference Tool

Haversack will be expanded to a full player reference tool. This phase will include full spellbook reference and spellcasting, bestiary reference, and allow for the creation and tracking of other creatures and NPCs.

## Phase Three: Player Journal

By this phase, Haversack will have reached all its primary goals. Everything from Phase Three and beyond will build on top of this core.

In Phase Three, Haversack will become a campaign journaling tool. NPCs can be recorded and tracked, important notes can be kept and reviewed, combat can be entirely logged. The player inventory system will be expanded to allow for better searching, tagging, and timed reminders. An in-game calendar system will allow for longer effects and durations and be useful for logging.

## Phase Four: Dungeon Master's Tool

[to be written, basically add more tools for DMs like battle maps, initiative trackers, quick reference for player character sheets, worldbuilding notes, etc]

## Phase Five: Full Party Tool

[to be written, basically allows a DM and several players to network together and use the clients at the same time in the same game. DMs can reveal and send information like notes, maps, and npc details, and can send messages and info privately to other players. Also adds the Nudge system, which lets DMs discreetly update or reorder notes, inventory items, etc. to subtly remind players of important, but forgotten, information. This is where the program becomes a competitor to something like D&DBeyond]