package coursevaniadownloader

// Course struct is used for storing the course related information
type Course struct {
	Name     string `json:"name"`
	MimeType string `json:"mimeType"`
}

// CourseFolder struct for storing the course related info
type CourseFolder struct {
	Files []File `json:"files"`
}

// File struct for storing the individual file info
type File struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	MimeType     string `json:"mimeType"`
	ModifiedTime string `json:"modifiedTime"`
}

// {
// 	"files": [
// 	  {
// 		"id": "1fAeBqbbwFJNtBI6fk2Ilad_xvAkGRYB6",
// 		"name": "1. Introduction",
// 		"mimeType": "application/vnd.google-apps.folder",
// 		"modifiedTime": "2019-11-03T16:50:49.381Z"
// 	  },
// 	  {
// 		"id": "1uhEHU2vBibCXxO5WNCwuqznUNS7eEu9f",
// 		"name": "10. Master Project Redux 2",
// 		"mimeType": "application/vnd.google-apps.folder",
// 		"modifiedTime": "2019-11-03T16:50:57.485Z"
// 	  },
// 	  {
// 		"id": "1Q6hwr1gi1ozBH7ghfX8TTnk7V1q6_Y1f",
// 		"name": "11. Master Project Advanced Routing",
// 		"mimeType": "application/vnd.google-apps.folder",
// 		"modifiedTime": "2019-11-03T16:51:21.764Z"
// 	  },
// 	  {
// 		"id": "16qrd-XXWHWPXN_uNIwu9kc6U1aOjU5p4",
// 		"name": "12. Master Project State Normalization",
// 		"mimeType": "application/vnd.google-apps.folder",
// 		"modifiedTime": "2019-11-03T16:51:39.487Z"
// 	  },
// 	  {
// 		"id": "1RXNo19VGYsUQMYP_jUk-bsHzPMzdw2Q-",
// 		"name": "13. Master Project Stripe Payments Part 1",
// 		"mimeType": "application/vnd.google-apps.folder",
// 		"modifiedTime": "2019-11-03T16:51:59.696Z"
// 	  },
// 	  {
// 		"id": "1lD3oe1nOXWej9S6F7F2YAP9Se_7t1fLt",
// 		"name": "14. Master Project Deploying To Production",
// 		"mimeType": "application/vnd.google-apps.folder",
// 		"modifiedTime": "2019-11-03T16:52:24.001Z"
// 	  },
// 	  {
// 		"id": "1-vY1nK1TQxLD2HjB5whQ6d4AxnjDHVAj",
// 		"name": "15. Master Project CSS in JS - styled-components",
// 		"mimeType": "application/vnd.google-apps.folder",
// 		"modifiedTime": "2019-11-03T16:52:45.983Z"
// 	  },
// 	  {
// 		"id": "1Ps8FSFSCUIXDZkIh0ZhT4vVOOBgrrCJo",
// 		"name": "16. Master Project Advanced Redux + Firebase",
// 		"mimeType": "application/vnd.google-apps.folder",
// 		"modifiedTime": "2019-11-03T16:53:22.988Z"
// 	  },
// 	  {
// 		"id": "1jHAp9XzXEFnp1fKdZCX1cP3KDCXQ-wvC",
// 		"name": "17. Master Project HOC Patterns",
// 		"mimeType": "application/vnd.google-apps.folder",
// 		"modifiedTime": "2019-11-03T16:54:05.656Z"
// 	  },
// 	  {
// 		"id": "1knUjjUHqPu0oxm8Ca9dldS0Xl9Ozb7lo",
// 		"name": "18. Master Project  Asynchronous Redux",
// 		"mimeType": "application/vnd.google-apps.folder",
// 		"modifiedTime": "2019-11-03T16:54:27.664Z"
// 	  },
// 	  {
// 		"id": "1DPzk6Dbfn067VqBhyLSbCVlhRcTlgU0M",
// 		"name": "19. Master Project Container Pattern",
// 		"mimeType": "application/vnd.google-apps.folder",
// 		"modifiedTime": "2019-11-03T16:54:59.739Z"
// 	  },
// 	  {
// 		"id": "1ATNTrQI7K-ADuPYy_maMwoHSWuagbiOw",
// 		"name": "2. React Key Concepts",
// 		"mimeType": "application/vnd.google-apps.folder",
// 		"modifiedTime": "2019-11-03T16:55:10.707Z"
// 	  },
// 	  {
// 		"id": "1_4q4Xsppk9VkOp0Br9fz6wOVQ5SV9oGd",
// 		"name": "20. Master Project Redux-Saga",
// 		"mimeType": "application/vnd.google-apps.folder",
// 		"modifiedTime": "2019-11-03T16:55:36.431Z"
// 	  },
// 	  {
// 		"id": "1MugGMWF66q9QxPGGkBd_Y6-3SYpQ1NzC",
// 		"name": "21. Master Project React Hooks",
// 		"mimeType": "application/vnd.google-apps.folder",
// 		"modifiedTime": "2019-11-03T16:57:06.399Z"
// 	  },
// 	  {
// 		"id": "1C_WT-rf31EkBTXyilu0qfLqXs1x-LBv4",
// 		"name": "22. Master Project Stripe Payments Part 2 - Backend",
// 		"mimeType": "application/vnd.google-apps.folder",
// 		"modifiedTime": "2019-11-03T16:58:02.729Z"
// 	  },
// 	  {
// 		"id": "17dvron8X08syBiDUlyLKrw0qgPL4szp0",
// 		"name": "23. Master Project Context API",
// 		"mimeType": "application/vnd.google-apps.folder",
// 		"modifiedTime": "2019-11-03T16:58:53.899Z"
// 	  },
// 	  {
// 		"id": "1Xq0Uu7ARrKG2M4sRBB3KQTk2Evg5WjlC",
// 		"name": "24. Master Project GraphQL + Apollo",
// 		"mimeType": "application/vnd.google-apps.folder",
// 		"modifiedTime": "2019-11-03T16:59:44.586Z"
// 	  },
// 	  {
// 		"id": "1bi0KOAIfPoZNufdv_ecakKhPJ1Z-LKDr",
// 		"name": "25. Master Project Mobile Support",
// 		"mimeType": "application/vnd.google-apps.folder",
// 		"modifiedTime": "2019-11-03T17:01:12.514Z"
// 	  },
// 	  {
// 		"id": "1h6XUTnUf3Rll2QQ0i8ygLznJgbY1prOx",
// 		"name": "26. Master Project React Performance",
// 		"mimeType": "application/vnd.google-apps.folder",
// 		"modifiedTime": "2019-11-03T17:01:36.673Z"
// 	  },
// 	  {
// 		"id": "1j517_OeUA3DRVPo5dbMxX9JRKG_xS35h",
// 		"name": "27. React Interview Questions + Advice",
// 		"mimeType": "application/vnd.google-apps.folder",
// 		"modifiedTime": "2019-11-03T17:02:59.263Z"
// 	  },
// 	  {
// 		"id": "1X8C-XQXo3kFANM6En36Sl_Z-ZXPQyml-",
// 		"name": "28. Bonus Progressive Web App",
// 		"mimeType": "application/vnd.google-apps.folder",
// 		"modifiedTime": "2019-11-03T17:03:14.906Z"
// 	  },
// 	  {
// 		"id": "1MYLKYq3AjiftevA8CnpwxL4j19ZkPAOV",
// 		"name": "29. Bonus Firebase Security",
// 		"mimeType": "application/vnd.google-apps.folder",
// 		"modifiedTime": "2019-11-03T17:04:02.662Z"
// 	  },
// 	  {
// 		"id": "1jGxPnKkArK4I26XQFBzM0iEOC2WdG1V4",
// 		"name": "3. React Basics",
// 		"mimeType": "application/vnd.google-apps.folder",
// 		"modifiedTime": "2019-11-03T17:04:28.933Z"
// 	  },
// 	  {
// 		"id": "1MixORydhFXZOnLdHB0ZU1ZtL-MqMtdli",
// 		"name": "30. Bonus Testing",
// 		"mimeType": "application/vnd.google-apps.folder",
// 		"modifiedTime": "2019-11-03T17:07:42.709Z"
// 	  },
// 	  {
// 		"id": "135efGH5sOu9VjbQhbLWTzxZbWVJOLBy1",
// 		"name": "31. Bonus Webpack + Babel",
// 		"mimeType": "application/vnd.google-apps.folder",
// 		"modifiedTime": "2019-11-03T17:10:06.276Z"
// 	  },
// 	  {
// 		"id": "1MR9n9VO-0inJh8YApKDonR2Cd3n01j6Z",
// 		"name": "32. Bonus Build a GatsbyJS Blog",
// 		"mimeType": "application/vnd.google-apps.folder",
// 		"modifiedTime": "2019-11-03T17:10:29.362Z"
// 	  },
// 	  {
// 		"id": "1xIICjzZD8dMSNcUFTSEeWbtKoUyuvv2A",
// 		"name": "33. Appendix 1 Key Developer Concepts",
// 		"mimeType": "application/vnd.google-apps.folder",
// 		"modifiedTime": "2019-11-03T17:11:53.886Z"
// 	  },
// 	  {
// 		"id": "1_x-dXDcvI_sghzAtQRw9GK4BXp7Ctxvh",
// 		"name": "34. Open Source Projects",
// 		"mimeType": "application/vnd.google-apps.folder",
// 		"modifiedTime": "2019-11-03T17:12:53.917Z"
// 	  },
// 	  {
// 		"id": "1ksZlq6G2NuZfm3PEa0-kqSPICg_10JdL",
// 		"name": "35. AMA + Bonus",
// 		"mimeType": "application/vnd.google-apps.folder",
// 		"modifiedTime": "2019-11-03T17:13:11.825Z"
// 	  },
// 	  {
// 		"id": "14eij0wC7vZJWdHWrAGgVs7BHySaSn4XU",
// 		"name": "36. Extras",
// 		"mimeType": "application/vnd.google-apps.folder",
// 		"modifiedTime": "2019-11-03T17:13:47.112Z"
// 	  },
// 	  {
// 		"id": "13Majumkw1ksnDjXr67U2GYGUqSA836s_",
// 		"name": "4. Master Project Setting Up E-commerce Project",
// 		"mimeType": "application/vnd.google-apps.folder",
// 		"modifiedTime": "2019-11-03T17:13:48.658Z"
// 	  },
// 	  {
// 		"id": "1Idj2EuCxN1i--XGMGNuAYqVwECoMcnsW",
// 		"name": "5. Master Project React Router and Routing",
// 		"mimeType": "application/vnd.google-apps.folder",
// 		"modifiedTime": "2019-11-03T17:14:47.455Z"
// 	  },
// 	  {
// 		"id": "128Mn9Ie3rOOt5sHrvLvXy9qEbMIcXAKz",
// 		"name": "6. Master Project Forms + Components",
// 		"mimeType": "application/vnd.google-apps.folder",
// 		"modifiedTime": "2019-11-03T17:15:14.771Z"
// 	  },
// 	  {
// 		"id": "16v_OoJ5OdhuuezMX4rzsLp02nJPM3e5d",
// 		"name": "7. Master Project Firebase + User Authentication",
// 		"mimeType": "application/vnd.google-apps.folder",
// 		"modifiedTime": "2019-11-03T17:16:10.737Z"
// 	  },
// 	  {
// 		"id": "1ernj__CTSRSVuVL3AVdHGgrgxjF39RlI",
// 		"name": "8. Master Project Redux 1",
// 		"mimeType": "application/vnd.google-apps.folder",
// 		"modifiedTime": "2019-11-03T17:17:52.593Z"
// 	  },
// 	  {
// 		"id": "1unZFtTNsZfxpEFXqUKCYvgsg3YOxA8sH",
// 		"name": "9. Master Project Session Storage + Persistence",
// 		"mimeType": "application/vnd.google-apps.folder",
// 		"modifiedTime": "2019-11-03T17:21:04.990Z"
// 	  }
// 	]
//   }
