// Copyright ©2013 The bíogo Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package entrez

import (
	"github.com/jjti/ncbi/entrez/link"
)

// <!--
//                 This is the Current DTD for Entrez eLink
// $Id: eLink_101123.dtd 349314 2012-01-09 23:26:00Z fialkov $
// -->
// <!-- ================================================================= -->
//
// <!ELEMENT	ERROR			(#PCDATA)>	<!-- .+ -->
// <!ELEMENT	Info			(#PCDATA)>	<!-- .+ -->
//
// <!ELEMENT	Id				(#PCDATA)>	<!-- \d+ -->
// <!ATTLIST	Id
// 			HasLinkOut  (Y|N)	#IMPLIED
// 			HasNeighbor (Y|N)	#IMPLIED
// 			>
//
// <!ELEMENT	Score			(#PCDATA)>	<!-- \d+ -->
// <!ELEMENT	DbFrom			(#PCDATA)>	<!-- \S+ -->
// <!ELEMENT	DbTo			(#PCDATA)>	<!-- \S+ -->
// <!ELEMENT	LinkName		(#PCDATA)>	<!-- \S+ -->
// <!ELEMENT	WebEnv			(#PCDATA)>	<!-- \S+ -->
// <!ELEMENT	MenuTag			(#PCDATA)>	<!-- \S+ -->
// <!ELEMENT	HtmlTag			(#PCDATA)>	<!-- \S+ -->
// <!ELEMENT	Priority		(#PCDATA)>	<!-- \S+ -->
//
// <!ELEMENT	IdList		(Id*)>
//
// <!-- cmd=neighbor -->
// <!ELEMENT	Link		(Id, Score?)>
// <!ELEMENT	QueryKey		(#PCDATA)>
//
// <!ELEMENT	LinkSetDb	(DbTo, LinkName, (Link*|Info), ERROR?)>
// <!ELEMENT	LinkSetDbHistory	(DbTo, LinkName, (QueryKey|Info), ERROR?)>
//
// <!-- cmd=llinks -->
//
// <!ELEMENT	Url			    (#PCDATA)>	<!-- \S+ -->
// <!ATTLIST	Url			LNG     (DA|DE|EN|EL|ES|FR|IT|IW|JA|NL|NO|RU|SV|ZH)     "EN">
//
// <!ELEMENT	IconUrl			(#PCDATA)>	<!-- \S+ -->
// <!ATTLIST	IconUrl		LNG     (DA|DE|EN|EL|ES|FR|IT|IW|JA|NL|NO|RU|SV|ZH)     "EN">
//
// <!ELEMENT	SubjectType		(#PCDATA)>	<!-- .+ -->
// <!ELEMENT	Category		(#PCDATA)>	<!-- .+ -->
// <!ELEMENT	Attribute		(#PCDATA)>	<!-- .+ -->
// <!--ELEMENT	LinkName		(#PCDATA)-->	<!--defined in neighbor section--><!-- \S+ -->
// <!ELEMENT	Name			(#PCDATA)>	<!-- .+ -->
// <!ELEMENT	NameAbbr		(#PCDATA)>	<!-- \S+ -->
// <!ELEMENT	SubProvider		(#PCDATA)>
//
// <!ELEMENT   FirstChar		(#PCDATA)>
//
// <!ELEMENT	Provider (
// 				Name,
// 				NameAbbr,
// 				Id,
// 				Url,
// 				IconUrl?
// 			)>
//
// <!ELEMENT	ObjUrl	(
// 				Url,
// 				IconUrl?,
// 				LinkName?,
//              SubjectType*,
// 				Category*,
//              Attribute*,
//              Provider,
//              SubProvider?
// 			)>
//
// <!ELEMENT	IdUrlSet	(Id,(ObjUrl+|Info))>
//
// <!ELEMENT	FirstChars  (FirstChar*)>
//
// <!ELEMENT	LinkInfo	(DbTo, LinkName, MenuTag?, HtmlTag?, Url?, Priority)>
// <!ELEMENT	IdLinkSet	(Id, LinkInfo*)>
// <!ELEMENT	IdUrlList	(IdUrlSet* | FirstChars*)>
//
// <!-- cmd=ncheck & lcheck & acheck -->
// <!ELEMENT	IdCheckList	((Id|IdLinkSet)*,ERROR?)>
//
// <!-- Common -->
// <!ELEMENT	LinkSet		(DbFrom,
// 				((IdList?, ((ERROR?, LinkSetDb)*  |  (LinkSetDbHistory*, WebEnv))) | IdUrlList | IdCheckList | ERROR), ERROR?
// 				)>
//
// <!ELEMENT	eLinkResult	(LinkSet*, ERROR?)>

// A Link holds the deserialised results of an ELink request.
type Link struct {
	LinkSets []link.LinkSet `xml:"LinkSet"`
	Err      *string        `xml:"ERROR"`
}
