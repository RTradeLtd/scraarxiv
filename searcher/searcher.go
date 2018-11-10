// Package searcher is used to search for research papers on arxiv
package searcher

import (
	"context"
	"fmt"
	"strings"

	arxiv "github.com/orijtech/arxiv/v1"
)

const (
	basePDFURL = "https://arxiv.org/pdf"
)

var (
	categories = []arxiv.Category{
		arxiv.StatisticsApplications,
		arxiv.StatisticsComputation,
		arxiv.StatisticsMachineLearning,
		arxiv.StatisticsMethodology,
		arxiv.StatisticsTheory,
		arxiv.QuantitativeBiologyBiomolecules,
		arxiv.QuantitativeBiologyCellBehavior,
		arxiv.QuantitativeBiologyGenomics,
		arxiv.QuantitativeBiologyMolecularNetworks,
		arxiv.QuantitativeBiologyNeuronsAndCognition,
		arxiv.QuantitativeBiologyOther,
		arxiv.QuantitativeBiologyPopulationsAndEvolution,
		arxiv.QuantitativeBiologyQuantitativeMethods,
		arxiv.QuantitativeBiologySubcellularProcesses,
		arxiv.QuantitativeBiologyTissuesAndOrgans,
		arxiv.CSArchitecture,
		arxiv.CSArtificialIntelligence,
		arxiv.CSComputationAndLanguage,
		arxiv.ComputationalComplexity,
		arxiv.ComputationalEngineeringFinanceAndScience,
		arxiv.CSComputationalGeometry,
		arxiv.CSGameTheory,
		arxiv.ComputerVisionAndPatternRecognition,
		arxiv.ComputersAndSociety,
		arxiv.CryptographyAndSecurity,
		arxiv.Databases,
		arxiv.DigitalLibraries,
		arxiv.DiscreteMathematics,
		arxiv.DistributedParallelAndClusterComputing,
		arxiv.CSGeneralLiterature,
		arxiv.CSGraphics,
		arxiv.HumanComputerInteraction,
		arxiv.CSInformationRetrieval,
		arxiv.CSInformationTheory,
		arxiv.CSLearning,
		arxiv.CSLogic,
		arxiv.CSMathematicalSoftware,
		arxiv.MultiagentSystems,
		arxiv.CSMultimedia,
		arxiv.NetworkAndInternetArchitecture,
		arxiv.NeuralAndEvolutionaryComputing,
		arxiv.CSNumericalAnalysis,
		arxiv.OperatingSystems,
		arxiv.CSOther,
		arxiv.CSPerformance,
		arxiv.ProgrammingLanguages,
		arxiv.CSRobotics,
		arxiv.SoftwareEngineering,
		arxiv.CSSound,
		arxiv.SymbolicComputation,
		arxiv.NonLinearSciencesAdaptationAndSelfOrganizingSystemsCategory,
		arxiv.NonLinearSciencesCellularAutomataAndLatticeGases,
		arxiv.NonLinearSciencesChaoticDynamics,
		arxiv.ExactlySolvableAndIntegrableSytems,
		arxiv.PatternFormationAndSolutions,
		arxiv.AlgebraicGeometry,
		arxiv.AlgebraicTopology,
		arxiv.AnalysisOfPDEs,
		arxiv.CategoryTheory,
		arxiv.ClassicalAnalysisAndODEs,
		arxiv.Combinatorics,
		arxiv.CommutativeAlgebra,
		arxiv.ComplexVariables,
		arxiv.DifferentialGeometry,
		arxiv.DynamicalSystems,
		arxiv.FunctionalAnalysis,
		arxiv.GeneralMathematics,
		arxiv.GeneralTopology,
		arxiv.GeometricTopology,
		arxiv.GroupTheory,
		arxiv.MathsHistoryAndOverview,
		arxiv.MathsInformationTheory,
		arxiv.KTheoryAndHomology,
		arxiv.MathsLogic,
		arxiv.MathsMathematicalPhysics,
		arxiv.MetricGeometry,
		arxiv.NumberTheory,
		arxiv.MathsNumericalAnalysis,
		arxiv.OperatorAlgebras,
		arxiv.MathsOptimizationAndControl,
		arxiv.Probability,
		arxiv.QuantumAlgebra,
		arxiv.RepresentationTheory,
		arxiv.RingsAndAlgebra,
		arxiv.MathsSpectralTheory,
		arxiv.MathsStatics,
		arxiv.SymplecticGeometry,
		arxiv.Astrophysics,
		arxiv.PhysicsDisorderedSystemsAndNeuralNetworks,
		arxiv.PhysicsMesoscopicSystemsAndQuantumHallEffect,
		arxiv.PhysicsMaterialsScience,
		arxiv.PhysicsOther,
		arxiv.PhysicsSoftCondensedMatter,
		arxiv.PhysicsStatisticalMechanics,
		arxiv.PhysicsStronglyCorrelatedElectrons,
		arxiv.PhysicsSuperconductivity,
		arxiv.GeneralRelativityAndQuantumCosmology,
		arxiv.HighEneryPhysicsExperiment,
		arxiv.HighEneryPhysicsLattice,
		arxiv.HighEneryPhysicsPhenomenology,
		arxiv.HighEneryPhysicsTheory,
		arxiv.MathematicalPhysics,
		arxiv.NuclearExperiment,
		arxiv.NuclearTheory,
		arxiv.AcceleratorPhysics,
		arxiv.AtmoshpericAndOceanicPhysics,
		arxiv.AtomicPhysics,
		arxiv.AtomicAndMolecularClusters,
		arxiv.BiologicalPhysics,
		arxiv.ChemicalPhysics,
		arxiv.ClassicalPhysics,
		arxiv.ComputationalPhysics,
		arxiv.DataAnalysisStatisticsAndProbability,
		arxiv.FluidDynamics,
		arxiv.GeneralPhysics,
		arxiv.Geophysics,
		arxiv.HistoryOfPhysics,
		arxiv.InstrumentationAndDetectors,
		arxiv.MedicalPhysics,
		arxiv.Optics,
		arxiv.PhysicsEducation,
		arxiv.PhysicsAndSociety,
		arxiv.PlasmaPhysics,
		arxiv.PopularPhysics,
		arxiv.SpacePhysics,
		arxiv.QuantumPhysics,
	}
)

// Search is used to perform a search against arxiv
func Search(term string, maxPageNumbers int64, maxCategories int) (map[string][]string, error) {
	//var urlsToScrape []string
	urlsToScrape := make(map[string][]string)

	for i, v := range categories {
		fmt.Println("fetching urls for category ", v)
		if maxCategories != 0 && i > maxCategories {
			break
		}
		// construct our query and generate a channel to receive data one
		responseChannel, cancel, err := arxiv.Search(
			context.Background(),
			&arxiv.Query{
				//Terms:         term,
				MaxPageNumber: maxPageNumbers,
				Filters: []*arxiv.Filter{
					{Op: arxiv.OpOR,
						Fields: []*arxiv.Field{
							{Category: v},
						},
					},
				},
			})
		if err != nil {
			return nil, err
		}
		for page := range responseChannel {
			// if this page had an error, skip it
			if err = page.Err; err != nil {
				fmt.Printf("error occured: %s\n", err)
				cancel()
				continue
			}
			for _, entry := range page.Feed.Entry {
				urlsToScrape[fmt.Sprintf("%s", v)] = append(urlsToScrape[fmt.Sprintf("%s", v)], entry.ID)
			}
		}
		cancel()
	}
	return urlsToScrape, nil
}

// ExtractPDFURLs is used to take an arxiv paper url, and get its pdf download equivalent
func ExtractPDFURLs(urls []string) []string {
	var pdfURLs []string
	for _, v := range urls {
		split := strings.Split(v, "/")
		url := fmt.Sprintf("%s/%s", basePDFURL, split[len(split)-1])
		pdfURLs = append(pdfURLs, url)
	}
	return pdfURLs
}
