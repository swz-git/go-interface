package gointerface

import (
	flat "github.com/swz-git/go-interface/flat"
)

//go:generate go run ./codegen/main.go

// Generated files, like the latter part of this file are
// intentionally not in .gitignore because of the go generate
// not being ran when you import a go package
//
// @generated
type AirState flat.AirState
type BallBouncinessOption flat.BallBouncinessOption
type BallInfo flat.BallInfoT
type BallMaxSpeedOption flat.BallMaxSpeedOption
type BallPrediction flat.BallPredictionT
type BallSizeOption flat.BallSizeOption
type BallTypeOption flat.BallTypeOption
type BallWeightOption flat.BallWeightOption
type Bool flat.BoolT
type BoostOption flat.BoostOption
type BoostPad flat.BoostPadT
type BoostPadState flat.BoostPadStateT
type BoostStrengthOption flat.BoostStrengthOption
type BoxShape flat.BoxShapeT
type CollisionShape flat.CollisionShape
type Color flat.ColorT
type ConsoleCommand flat.ConsoleCommandT
type ControllerState flat.ControllerStateT
type CylinderShape flat.CylinderShapeT
type DemolishOption flat.DemolishOption
type DesiredBallState flat.DesiredBallStateT
type DesiredBoostState flat.DesiredBoostStateT
type DesiredCarState flat.DesiredCarStateT
type DesiredGameInfoState flat.DesiredGameInfoStateT
type DesiredGameState flat.DesiredGameStateT
type DesiredPhysics flat.DesiredPhysicsT
type ExistingMatchBehavior flat.ExistingMatchBehavior
type FieldInfo flat.FieldInfoT
type Float flat.FloatT
type GameInfo flat.GameInfoT
type GameMessage flat.GameMessage
type GameMessageWrapper flat.GameMessageWrapperT
type GameMode flat.GameMode
type GameSpeedOption flat.GameSpeedOption
type GameStateType flat.GameStateType
type GameTickPacket flat.GameTickPacketT
type GoalInfo flat.GoalInfoT
type GravityOption flat.GravityOption
type Human flat.HumanT
type Launcher flat.Launcher
type Line3D flat.Line3DT
type LoadoutPaint flat.LoadoutPaintT
type MatchComm flat.MatchCommT
type MatchLength flat.MatchLength
type MatchSettings flat.MatchSettingsT
type MaxScore flat.MaxScore
type MessagePacket flat.MessagePacketT
type MutatorSettings flat.MutatorSettingsT
type OvertimeOption flat.OvertimeOption
type PartyMember flat.PartyMemberT
type Physics flat.PhysicsT
type PlayerClass flat.PlayerClass
type PlayerConfiguration flat.PlayerConfigurationT
type PlayerInfo flat.PlayerInfoT
type PlayerInput flat.PlayerInputT
type PlayerInputChange flat.PlayerInputChangeT
type PlayerLoadout flat.PlayerLoadoutT
type PlayerSpectate flat.PlayerSpectateT
type PlayerStatEvent flat.PlayerStatEventT
type PolyLine3D flat.PolyLine3DT
type PredictionSlice flat.PredictionSliceT
type Psyonix flat.PsyonixT
type RLBot flat.RLBotT
type ReadyMessage flat.ReadyMessageT
type RemoveRenderGroup flat.RemoveRenderGroupT
type RenderGroup flat.RenderGroupT
type RenderMessage flat.RenderMessageT
type RenderType flat.RenderType
type RespawnTimeOption flat.RespawnTimeOption
type Rotator flat.RotatorT
type RotatorPartial flat.RotatorPartialT
type RumbleOption flat.RumbleOption
type ScoreInfo flat.ScoreInfoT
type ScriptConfiguration flat.ScriptConfigurationT
type SeriesLengthOption flat.SeriesLengthOption
type SphereShape flat.SphereShapeT
type StartCommand flat.StartCommandT
type StopCommand flat.StopCommandT
type String2D flat.String2DT
type String3D flat.String3DT
type TeamInfo flat.TeamInfoT
type TextHAlign flat.TextHAlign
type TextVAlign flat.TextVAlign
type Touch flat.TouchT
type Vector3 flat.Vector3T
type Vector3Partial flat.Vector3PartialT
