package main

import (
	"fmt"
	"log"
	"math/cmplx"

	"thegame/pb"
)

type Hero struct {
	Entity

	score         int
	experience    int
	skillPoints   int
	level         int
	abilityLevels [NAbilities]int
	orientation   float64
	cooldown      int
	id            int

	controls   *pb.Controls
	UpdateChan chan *pb.GameState
}

func (h *Hero) ToProto() *pb.Hero {
	levels := make([]int32, 8)
	values := make([]int32, 8)
	for i, level := range h.abilityLevels {
		levels[i] = int32(level)
		values[i] = int32(AbilityValues[i][level])
	}
	return &pb.Hero{
		Entity:              EntityToProto(h),
		AbilityLevels:       levels,
		AbilityValues:       values,
		Orientation:         h.orientation,
		Level:               int32(h.level),
		Score:               int32(h.score),
		Experience:          int32(h.experience),
		ExperienceToLevelUp: int32(experienceToLevelUp(h.level)),
	}
}

func NewHero(id int) *Hero {
	return &Hero{
		level:      1,
		id:         id,
		UpdateChan: make(chan *pb.GameState, 1),
	}
}

func (h *Hero) ID() int {
	return h.id
}

func (h *Hero) String() string {
	return fmt.Sprintf("Hero#%d", h.id)
}

func (h *Hero) Friction() float64 {
	return 0.1
}

func (h *Hero) IsBounded() bool {
	return true
}

func (h *Hero) MaxSpeed() float64 {
	return float64(h.ability(MovementSpeed))
}

func (h *Hero) Radius() float64 {
	return 30
}

func (h *Hero) Team() int {
	return h.id
}

func (h *Hero) ability(a int) int {
	return AbilityValues[a][h.abilityLevels[a]]
}

func (h *Hero) BodyDamage() int {
	return h.ability(BodyDamage)
}

func (h *Hero) AcquireExperience(e int) {
	h.score += e
	h.experience += e
	for h.experience >= experienceToLevelUp(h.level) {
		h.experience -= experienceToLevelUp(h.level)
		h.level++
		h.skillPoints++
	}
}

func (h *Hero) RewardingExperience() int {
	return h.score / 2
}

func (h *Hero) MaxHealth() int {
	return h.ability(MaxHealth)
}

// Shoot creates a bullet from the hero's stats
// After shooting, remember to assign a bullet id
func (h *Hero) Shoot() *Bullet {
	return &Bullet{
		Entity: Entity{
			health:   h.ability(BulletPenetration),
			position: h.position + h.velocity + cmplx.Rect(h.Radius(), h.orientation),
			velocity: cmplx.Rect(float64(h.ability(BulletSpeed)), h.orientation),
			visible:  true,
		},
		owner:   h,
		timeout: 100,
	}
}

// Action performs the action for the current tick on the given arena
func (h *Hero) Action(a *Arena) {
	if h.controls == nil {
		return // TODO
	}
	h.orientation = h.controls.ShootDirection
	if h.cooldown > 0 {
		h.cooldown--
	} else if h.controls.Shoot {
		bullet := h.Shoot()
		a.bulletCounter++
		bullet.id = a.bulletCounter
		a.bullets = append(a.bullets, bullet)
		h.cooldown = h.ability(Reload)
	}
	if h.health > 0 {
		h.health += h.ability(HealthRegen)
		if h.health > h.ability(MaxHealth) {
			h.health = h.ability(MaxHealth)
		}
	}
	for _, skill := range h.controls.LevelUp {
		if h.skillPoints <= 0 {
			break
		}
		if h.abilityLevels[skill] < 8 {
			h.abilityLevels[skill]++
			h.skillPoints--
		}
	}
	if h.controls.Accelerate {
		h.velocity += cmplx.Rect(0.6, h.controls.AccelerationDirection)
	}
}

func (h *Hero) Spawn() {
	experience := h.score / 2
	h.score = 0
	h.experience = 0
	h.health = h.MaxHealth()
	h.visible = true
	h.AcquireExperience(experience)
	h.position = RandomPosition()
	h.velocity = 0
	log.Println(h, "spawned at", h.position)
}
