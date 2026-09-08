### Features included
- ...

### Features to be done
- [ ] interior space
  - [ ] tiles
  - [ ] walls/major colliders
  - [ ] furniture/minor colliders
  - [ ] ability to place items on top of objects
- [ ] player object
  - [ ] movement
  - [ ] direction recognition (4 ways)
    - [ ] interaction based on direction
    - [ ] recognising items and activators
- [ ] door/window objects
- [ ] movement - expansion
  - [ ] travel between interiors (doors)
  - [ ] moving the camera with player (test on corridor)
  - [ ] respecting interior space and its interactions alongside movement
- [ ] passing time (showcased by window state)
- [ ] bed (allows passing time)
- [ ] UI
  - [ ] minimal UI for activators (e.g. bed, dialogue)
  - [ ] minimal UI for inventory (ability to pick up items?)
  - [ ] minimal UI for context (e.g. time, health, etc.)
- [ ] NPCs
  - [ ] schedules
  - [ ] pathing (if collision isn't enough)
  - [ ] dialogues
- [ ] minimal world (4 rooms + corridor), with items, NPCs, their dialogues and schedules
- [ ] center the interior map, so no `empty` tiles need to be marked
---
<!-- scripts -->
- [ ] scripts
  - [ ] Lua/Python interfacing
  - [ ] done through predefined functions
    - [ ] `runScript()` to run another script (chaining)
    - [ ] `initialiseVariable()`
    - [ ] `set/getVariable()`
    - [ ] figuring out whether NPC and player script should be
      - separated (`playerMove()`/`npcMove()`)
      - abstract (`entityMove(player)`)
    - [ ] `openMenu()` (e.g. for dialogues, GUIs)
  - [ ] always-active scripts
    - [ ] set it programmatically or by listing files?
    - [ ] event handling/subscription
      - [ ] check performance issues by global events check
- [ ] activators
  - [ ] running scripts
  - [ ] teleporting by door activators
---
<!-- customisation -->
- [ ] customisable GUI
- [ ] check if you can do some mechanics via scripts, e.g. crafting
---
<!-- exterior space -->
- [ ] background (movable?)
- [ ] weather
  - [ ] interaction with interiors? (sound)
- [ ] regions? seasons?
---
<!-- finals -->
- [ ] background music & interaction/movement sounds
- [ ] finishing touches
  - [ ] menu with `new game`/`load`/`settings` options
  - [ ] options for resolution and volume
  - [ ] save/load system
