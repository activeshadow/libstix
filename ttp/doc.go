// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

// Package ttp provides support for the STIX TTP idiom.
//
// TTPs are representations of the behavior or modus operandi of cyber
// adversaries. It is a term taken from the traditional military sphere and is
// used to characterize what an adversary does and how they do it in increasing
// levels of detail. For instance, to give a simple example, a tactic may be to
// use malware to steal credit card credentials. A related technique (at a lower
// level of detail) may be to send targeted emails to potential victims, which
// have documents attached containing malicious code which executes upon
// opening, captures credit card information from keystrokes, and uses http to
// communicate with a command and control server to transfer information. A
// related procedure (at a lower level of detail) may be to perform open source
// research to identify potentially gullible individuals, craft a convincing
// socially engineered email and document, create malware/exploit that will
// bypass current antivirus detection, establish a command and control server by
// registering a domain called mychasebank.org, and send mail to victims from a
// Gmail account called accounts-mychasebank@gmail.com.
//
// TTPs consist of the specific adversary behavior (attack patterns, malware,
// exploits) exhibited, resources leveraged (tools, infrastructure, personas),
// information on the victims targeted (who, what or where), relevant
// ExploitTargets being targeted, intended effects, relevant kill chain phases,
// handling guidance, source of the TTP information, etc.
//
// TTPs play a central role in cyber threat information and cyber threat
// intelligence. They are relevant for Indicators, Incidents, Campaigns, and
// ThreatActors. In addition, they hold a close relationship with ExploitTargets
//  that characterize the specific targets that the TTPs seek to exploit.
package ttp
