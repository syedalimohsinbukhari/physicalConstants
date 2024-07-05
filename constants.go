package physicalConstants

import (
	"fmt"
	"math"
)

const PI_ = math.Pi
const OHM_ = "\u03A9"

var pow = math.Pow

type physicalConstant struct {
	Value float64
	Unit  string
}

func (p physicalConstant) Display() {
	fmt.Println(p.Value, p.Unit)
}

func CustomPhysicalConstant(Value float64, Unit string) physicalConstant {
	return physicalConstant{Value: Value, Unit: Unit}
}

var ElectronMass = physicalConstant{
	Value: 9.1093837139e-31,
	Unit:  "kg",
}

var ProtonMass = physicalConstant{
	Value: 1.67262192595e-27,
	Unit:  "kg",
}

var NeutronMass = physicalConstant{
	Value: 1.67492750056e-27,
	Unit:  "kg",
}

var MuonMass = physicalConstant{
	Value: 1.883531627e-28,
	Unit:  "kg",
}

var TopQuarkMass = physicalConstant{
	Value: 3.0784e-25,
	Unit:  "kg",
}

var ProtonElectronMassRatio = physicalConstant{
	Value: ProtonMass.Value / ElectronMass.Value,
	Unit:  "",
}

var TauMass = physicalConstant{
	Value: 3.16754e-27,
	Unit:  "kg",
}

var SpeedOfLight = physicalConstant{
	Value: 299792458,
	Unit:  "m/s",
}

var PlanckConstant = physicalConstant{
	Value: 6.62607015e-34,
	Unit:  "J.s",
}

var ReducedPlanckConstant = physicalConstant{
	Value: PlanckConstant.Value / (2 * PI_),
	Unit:  PlanckConstant.Unit,
}

var NewtonianConstantOfGravitation = physicalConstant{
	Value: 6.67430e-11,
	Unit:  "m^3/kg/s^2",
}

var GravitationalConstant = physicalConstant{
	Value: NewtonianConstantOfGravitation.Value,
	Unit:  NewtonianConstantOfGravitation.Unit,
}

var CosmologicalConstant = physicalConstant{
	Value: 1.089e-52,
	Unit:  "1/m^2",
}

var WienConstantWavelength = physicalConstant{
	Value: 2.897771955e-3,
	Unit:  "m.K",
}

var WienConstantFrequency = physicalConstant{
	Value: 5.878925757e10,
	Unit:  "Hz/K",
}

var WienConstantEntropy = physicalConstant{
	Value: 3.002916077e-3,
	Unit:  "m.K",
}

var ElementaryCharge = physicalConstant{
	Value: 1.602176634e-19,
	Unit:  "C",
}

var BoltzmannConstant = physicalConstant{
	Value: 1.380649e-23,
	Unit:  "J/K",
}

var StefanBoltzmannConstant = physicalConstant{
	Value: stefanBoltzmannConstant(),
	Unit:  "W/m^2/K^4",
}

var FirstRadiationConstant = physicalConstant{
	Value: 2 * PI_ * PlanckConstant.Value * pow(SpeedOfLight.Value, 2),
	Unit:  "W.m^2",
}

var FirstRadiationConstantSpectralRadiance = physicalConstant{
	Value: 2 * PlanckConstant.Value * pow(SpeedOfLight.Value, 2),
	Unit:  "W.m^2/sr",
}

var SecondRadiationConstant = physicalConstant{
	Value: PlanckConstant.Value * SpeedOfLight.Value * pow(BoltzmannConstant.Value, -1),
	Unit:  "m.K",
}

var FineStructureConstant = physicalConstant{
	Value: 1 / 137.,
	Unit:  "",
}

var VacuumMagneticPermeability = physicalConstant{
	Value: vacuumMagneticPermeability(),
	Unit:  "N/A^2",
}

var VacuumElectricPermeability = physicalConstant{
	Value: vacuumElectricPermeability(),
	Unit:  "F/m",
}

var ConductanceQuantum = physicalConstant{
	Value: 2 * pow(ElementaryCharge.Value, 2) / PlanckConstant.Value,
	Unit:  "S",
}

var InverseConductanceQuantum = physicalConstant{
	Value: 1 / ConductanceQuantum.Value,
	Unit:  OHM_,
}

var VonKlitzingConstant = physicalConstant{
	Value: PlanckConstant.Value / pow(ElementaryCharge.Value, 2),
	Unit:  OHM_,
}

var JosephsonConstant = physicalConstant{
	Value: (2 * ElementaryCharge.Value) / PlanckConstant.Value,
	Unit:  "Hz/V",
}

var MagneticFluxQuantum = physicalConstant{
	Value: 1 / JosephsonConstant.Value,
	Unit:  "V.s",
}

var CharacteristicImpedanceOfVacuum = physicalConstant{
	Value: characteristicImpedanceOfVacuum(),
	Unit:  OHM_,
}

var CarbonMolarMass = physicalConstant{
	Value: 12.0000000126e-3,
	Unit:  "kg/mol",
}

var AtomicMassUnit = physicalConstant{
	Value: 1.66053906892e-27,
	Unit:  "kg",
}

var AvogadroNumber = physicalConstant{
	Value: 6.02214076e23,
	Unit:  "1/mol",
}

var MolarPlanckConstant = physicalConstant{
	Value: AvogadroNumber.Value * PlanckConstant.Value,
	Unit:  "J.s/mol",
}

var MolarGasConstant = physicalConstant{
	Value: AvogadroNumber.Value * BoltzmannConstant.Value,
	Unit:  "J/mol/K",
}

var FaradayConstant = physicalConstant{
	Value: AvogadroNumber.Value * ElementaryCharge.Value,
	Unit:  "C/mol",
}

var MolarMassConstant = physicalConstant{
	Value: 1.00000000105e-3,
	Unit:  "kg/mol",
}

var BohrRadius = physicalConstant{
	Value: bohrRadius(),
	Unit:  "m",
}

var ClassicalElectronRadius = physicalConstant{
	Value: classicalElectronRadius(),
	Unit:  "m",
}

var ThomsonCrossSection = physicalConstant{
	Value: thomsonCrossSection(),
	Unit:  "m^2",
}

var RydbergConstant = physicalConstant{
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
