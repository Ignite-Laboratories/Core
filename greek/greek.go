package greek

import "github.com/ignite-laboratories/core"

// Lowercase provides all the lowercase Greek characters.
var Lowercase = []core.GivenName{
	Alpha, Beta, Gamma, Delta, Epsilon, Zeta, Eta, Theta, Iota, Kappa, Lambda, Mu, Nu, Xi, Omicron, Pi, Rho, Sigma, SigmaFinal, Tau, Upsilon, Phi, Chi, Psi, Omega,
}

// Uppercase provides all the uppercase Greek characters.
var Uppercase = []core.GivenName{
	AlphaUpper, BetaUpper, GammaUpper, DeltaUpper, EpsilonUpper, ZetaUpper, EtaUpper, ThetaUpper, IotaUpper, KappaUpper, LambdaUpper, MuUpper, NuUpper, XiUpper, OmicronUpper, PiUpper, RhoUpper, SigmaUpper, TauUpper, UpsilonUpper, PhiUpper, ChiUpper, PsiUpper, OmegaUpper,
}

// Anycase provides all Greek characters, regardless of case.
var Anycase = []core.GivenName{
	Alpha, Beta, Gamma, Delta, Epsilon, Zeta, Eta, Theta, Iota, Kappa, Lambda, Mu, Nu, Xi, Omicron, Pi, Rho, Sigma, SigmaFinal, Tau, Upsilon, Phi, Chi, Psi, Omega,
	AlphaUpper, BetaUpper, GammaUpper, DeltaUpper, EpsilonUpper, ZetaUpper, EtaUpper, ThetaUpper, IotaUpper, KappaUpper, LambdaUpper, MuUpper, NuUpper, XiUpper, OmicronUpper, PiUpper, RhoUpper, SigmaUpper, TauUpper, UpsilonUpper, PhiUpper, ChiUpper, PsiUpper, OmegaUpper,
}

// AlphaUpper represents the Greek character "Α".
var AlphaUpper, _ = core.LookupName("Α")

// Alpha represents the Greek character "α".
var Alpha, _ = core.LookupName("α")

// BetaUpper represents the Greek character "Β".
var BetaUpper, _ = core.LookupName("Β")

// Beta represents the Greek character "β".
var Beta, _ = core.LookupName("β")

// GammaUpper represents the Greek character "Γ".
var GammaUpper, _ = core.LookupName("Γ")

// Gamma represents the Greek character "γ".
var Gamma, _ = core.LookupName("γ")

// DeltaUpper represents the Greek character "Δ".
var DeltaUpper, _ = core.LookupName("Δ")

// Delta represents the Greek character "δ".
var Delta, _ = core.LookupName("δ")

// EpsilonUpper represents the Greek character "Ε".
var EpsilonUpper, _ = core.LookupName("Ε")

// Epsilon represents the Greek character "ε".
var Epsilon, _ = core.LookupName("ε")

// ZetaUpper represents the Greek character "Ζ".
var ZetaUpper, _ = core.LookupName("Ζ")

// Zeta represents the Greek character "ζ".
var Zeta, _ = core.LookupName("ζ")

// EtaUpper represents the Greek character "Η".
var EtaUpper, _ = core.LookupName("Η")

// Eta represents the Greek character "Η".
var Eta, _ = core.LookupName("η")

// ThetaUpper represents the Greek character "Θ".
var ThetaUpper, _ = core.LookupName("Θ")

// Theta represents the Greek character "θ".
var Theta, _ = core.LookupName("θ")

// IotaUpper represents the Greek character "Ι".
var IotaUpper, _ = core.LookupName("Ι")

// Iota represents the Greek character "ι".
var Iota, _ = core.LookupName("ι")

// KappaUpper represents the Greek character "Κ".
var KappaUpper, _ = core.LookupName("Κ")

// Kappa represents the Greek character "κ".
var Kappa, _ = core.LookupName("κ")

// LambdaUpper represents the Greek character "Λ".
var LambdaUpper, _ = core.LookupName("Λ")

// Lambda represents the Greek character "λ".
var Lambda, _ = core.LookupName("λ")

// MuUpper represents the Greek character "Μ".
var MuUpper, _ = core.LookupName("Μ")

// Mu represents the Greek character "μ".
var Mu, _ = core.LookupName("μ")

// NuUpper represents the Greek character "Ν".
var NuUpper, _ = core.LookupName("Ν")

// Nu represents the Greek character "ν".
var Nu, _ = core.LookupName("ν")

// XiUpper represents the Greek character "Ξ".
var XiUpper, _ = core.LookupName("Ξ")

// Xi represents the Greek character "ξ".
var Xi, _ = core.LookupName("ξ")

// OmicronUpper represents the Greek character "Ο".
var OmicronUpper, _ = core.LookupName("Ο")

// Omicron represents the Greek character "ο".
var Omicron, _ = core.LookupName("ο")

// PiUpper represents the Greek character "Π".
var PiUpper, _ = core.LookupName("Π")

// Pi represents the Greek character "π".
var Pi, _ = core.LookupName("π")

// RhoUpper represents the Greek character "Ρ".
var RhoUpper, _ = core.LookupName("Ρ")

// Rho represents the Greek character "ρ".
var Rho, _ = core.LookupName("ρ")

// SigmaUpper represents the Greek character "Σ".
var SigmaUpper, _ = core.LookupName("Σ")

// Sigma represents the Greek character "σ".
//
// NOTE: This is the standard form, see SigmaFinal for the final form.
var Sigma, _ = core.LookupName("σ")

// SigmaFinal represents the Greek character "ς".
//
// NOTE: This is final form, see Sigma for the standard form.
var SigmaFinal, _ = core.LookupName("ς")

// TauUpper represents the Greek character "Τ".
var TauUpper, _ = core.LookupName("Τ")

// Tau represents the Greek character "τ".
var Tau, _ = core.LookupName("τ")

// UpsilonUpper represents the Greek character "Υ".
var UpsilonUpper, _ = core.LookupName("Υ")

// Upsilon represents the Greek character "υ".
var Upsilon, _ = core.LookupName("υ")

// PhiUpper represents the Greek character "Φ".
var PhiUpper, _ = core.LookupName("Φ")

// Phi represents the Greek character "φ".
var Phi, _ = core.LookupName("φ")

// ChiUpper represents the Greek character "Χ".
var ChiUpper, _ = core.LookupName("Χ")

// Chi represents the Greek character "χ".
var Chi, _ = core.LookupName("χ")

// PsiUpper represents the Greek character "Ψ".
var PsiUpper, _ = core.LookupName("Ψ")

// Psi represents the Greek character "ψ".
var Psi, _ = core.LookupName("ψ")

// OmegaUpper represents the Greek character "Ω".
var OmegaUpper, _ = core.LookupName("Ω")

// Omega represents the Greek character "ω".
var Omega, _ = core.LookupName("ω")
