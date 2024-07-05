package physicalConstants

import (
	"math"
)

const PI_ = math.Pi
const OHM_ = "\u03A9"

var pow = math.Pow

var ElectronMass = PhysicalConstant{
	Value: 9.1093837139e-31,
	Unit:  "kg",
}

var ProtonMass = PhysicalConstant{
	Value: 1.67262192595e-27,
	Unit:  "kg",
}

var NeutronMass = PhysicalConstant{
	Value: 1.67492750056e-27,
	Unit:  "kg",
}

var MuonMass = PhysicalConstant{
	Value: 1.883531627e-28,
	Unit:  "kg",
}

var TopQuarkMass = PhysicalConstant{
	Value: 3.0784e-25,
	Unit:  "kg",
}

var ProtonElectronMassRatio = PhysicalConstant{
	Value: ProtonMass.Value / ElectronMass.Value,
	Unit:  "",
}

var TauMass = PhysicalConstant{
	Value: 3.16754e-27,
	Unit:  "kg",
}

var SpeedOfLight = PhysicalConstant{
	Value: 299792458,
	Unit:  "m/s",
}

var PlanckConstant = PhysicalConstant{
	Value: 6.62607015e-34,
	Unit:  "J.s",
}

var ReducedPlanckConstant = PhysicalConstant{
	Value: PlanckConstant.Value / (2 * PI_),
	Unit:  PlanckConstant.Unit,
}

var NewtonianConstantOfGravitation = PhysicalConstant{
	Value: 6.67430e-11,
	Unit:  "m^3/kg/s^2",
}

var GravitationalConstant = PhysicalConstant{
	Value: NewtonianConstantOfGravitation.Value,
	Unit:  NewtonianConstantOfGravitation.Unit,
}

var CosmologicalConstant = PhysicalConstant{
	Value: 1.089e-52,
	Unit:  "1/m^2",
}

var WienConstantWavelength = PhysicalConstant{
	Value: 2.897771955e-3,
	Unit:  "m.K",
}

var WienConstantFrequency = PhysicalConstant{
	Value: 5.878925757e10,
	Unit:  "Hz/K",
}

var WienConstantEntropy = PhysicalConstant{
	Value: 3.002916077e-3,
	Unit:  "m.K",
}

var ElementaryCharge = PhysicalConstant{
	Value: 1.602176634e-19,
	Unit:  "C",
}

var BoltzmannConstant = PhysicalConstant{
	Value: 1.380649e-23,
	Unit:  "J/K",
}

var StefanBoltzmannConstant = PhysicalConstant{
	Value: stefanBoltzmannConstant(),
	Unit:  "W/m^2/K^4",
}

var FirstRadiationConstant = PhysicalConstant{
	Value: 2 * PI_ * PlanckConstant.Value * pow(SpeedOfLight.Value, 2),
	Unit:  "W.m^2",
}

var FirstRadiationConstantSpectralRadiance = PhysicalConstant{
	Value: 2 * PlanckConstant.Value * pow(SpeedOfLight.Value, 2),
	Unit:  "W.m^2/sr",
}

var SecondRadiationConstant = PhysicalConstant{
	Value: PlanckConstant.Value * SpeedOfLight.Value * pow(BoltzmannConstant.Value, -1),
	Unit:  "m.K",
}

var FineStructureConstant = PhysicalConstant{
	Value: 1 / 137.,
	Unit:  "",
}

var VacuumMagneticPermeability = PhysicalConstant{
	Value: vacuumMagneticPermeability(),
	Unit:  "N/A^2",
}

var VacuumElectricPermeability = PhysicalConstant{
	Value: vacuumElectricPermeability(),
	Unit:  "F/m",
}

var ConductanceQuantum = PhysicalConstant{
	Value: 2 * pow(ElementaryCharge.Value, 2) / PlanckConstant.Value,
	Unit:  "S",
}

var InverseConductanceQuantum = PhysicalConstant{
	Value: 1 / ConductanceQuantum.Value,
	Unit:  OHM_,
}

var VonKlitzingConstant = PhysicalConstant{
	Value: PlanckConstant.Value / pow(ElementaryCharge.Value, 2),
	Unit:  OHM_,
}

var JosephsonConstant = PhysicalConstant{
	Value: (2 * ElementaryCharge.Value) / PlanckConstant.Value,
	Unit:  "Hz/V",
}

var MagneticFluxQuantum = PhysicalConstant{
	Value: 1 / JosephsonConstant.Value,
	Unit:  "V.s",
}

var CharacteristicImpedanceOfVacuum = PhysicalConstant{
	Value: characteristicImpedanceOfVacuum(),
	Unit:  OHM_,
}

var CarbonMolarMass = PhysicalConstant{
	Value: 12.0000000126e-3,
	Unit:  "kg/mol",
}

var AtomicMassUnit = PhysicalConstant{
	Value: 1.66053906892e-27,
	Unit:  "kg",
}

var AvogadroNumber = PhysicalConstant{
	Value: 6.02214076e23,
	Unit:  "1/mol",
}

var MolarPlanckConstant = PhysicalConstant{
	Value: AvogadroNumber.Value * PlanckConstant.Value,
	Unit:  "J.s/mol",
}

var MolarGasConstant = PhysicalConstant{
	Value: AvogadroNumber.Value * BoltzmannConstant.Value,
	Unit:  "J/mol/K",
}

var FaradayConstant = PhysicalConstant{
	Value: AvogadroNumber.Value * ElementaryCharge.Value,
	Unit:  "C/mol",
}

var MolarMassConstant = PhysicalConstant{
	Value: 1.00000000105e-3,
	Unit:  "kg/mol",
}

var BohrRadius = PhysicalConstant{
	Value: bohrRadius(),
	Unit:  "m",
}

var ClassicalElectronRadius = PhysicalConstant{
	Value: classicalElectronRadius(),
	Unit:  "m",
}

var ThomsonCrossSection = PhysicalConstant{
	Value: thomsonCrossSection(),
	Unit:  "m^2",
}

var RydbergConstant = PhysicalConstant{
	Value: rydbergConstant(),
	Unit:  "1/m",
}

func rydbergConstant() float64 {
	num_ := pow(FineStructureConstant.Value, 2) * ElectronMass.Value * SpeedOfLight.Value
	den_ := 2 * PlanckConstant.Value

	return num_ / den_
}

func thomsonCrossSection() float64 {
	num_ := (8 * PI_) / 3
	den_ := pow(classicalElectronRadius(), 2)

	return num_ / den_
}

func classicalElectronRadius() float64 {
	num_ := FineStructureConstant.Value * ReducedPlanckConstant.Value
	den_ := ElectronMass.Value * SpeedOfLight.Value

	return num_ / den_
}

func bohrRadius() float64 {
	num_ := ReducedPlanckConstant.Value
	den_ := FineStructureConstant.Value * ElectronMass.Value * SpeedOfLight.Value

	return num_ / den_
}

func characteristicImpedanceOfVacuum() float64 {
	num_ := 4 * PI_ * FineStructureConstant.Value * ReducedPlanckConstant.Value
	den_ := pow(ElementaryCharge.Value, 2)

	return num_ / den_
}

func vacuumMagneticPermeability() float64 {
	num_ := 4 * PI_ * FineStructureConstant.Value * ReducedPlanckConstant.Value
	den_ := pow(ElementaryCharge.Value, 2) * SpeedOfLight.Value

	return num_ / den_
}

func vacuumElectricPermeability() float64 {
	num_ := 1 / VacuumMagneticPermeability.Value
	return num_ / pow(SpeedOfLight.Value, 2)
}

func stefanBoltzmannConstant() float64 {
	num_ := pow(PI_, 2) * pow(BoltzmannConstant.Value, 4)
	den_ := 60 * pow(ReducedPlanckConstant.Value, 3) * pow(SpeedOfLight.Value, 2)

	return num_ / den_
}
