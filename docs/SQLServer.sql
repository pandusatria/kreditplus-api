USE [KreditPlus]
GO
/****** Object:  Table [dbo].[tbl_employee]    Script Date: 12/7/2018 2:16:20 PM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[tbl_employee](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[userid] [int] NOT NULL,
	[firstname] [varchar](100) NOT NULL,
	[lastname] [varchar](100) NOT NULL,
	[jobtitle] [varchar](100) NULL,
	[salary] [decimal](10, 2) NULL,
	[created_date] [date] NOT NULL,
	[created_by] [varchar](100) NOT NULL,
	[updated_date] [date] NULL,
	[updated_by] [varchar](100) NULL,
 CONSTRAINT [PK_tbl_employee] PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]
GO
SET IDENTITY_INSERT [dbo].[tbl_employee] ON 

INSERT [dbo].[tbl_employee] ([id], [userid], [firstname], [lastname], [jobtitle], [salary], [created_date], [created_by], [updated_date], [updated_by]) VALUES (1, 1, N'Satria', N'Baja Hitam', N'Developer', CAST(6500000.00 AS Decimal(10, 2)), CAST(N'2018-11-28' AS Date), N'system', CAST(N'2018-12-06' AS Date), N'system')
INSERT [dbo].[tbl_employee] ([id], [userid], [firstname], [lastname], [jobtitle], [salary], [created_date], [created_by], [updated_date], [updated_by]) VALUES (2, 7, N'Lionel', N'Messi', N'Web Developer', CAST(5000000.00 AS Decimal(10, 2)), CAST(N'2018-12-05' AS Date), N'system', NULL, NULL)
SET IDENTITY_INSERT [dbo].[tbl_employee] OFF
/****** Object:  StoredProcedure [dbo].[Sp_GetAllEmployee]    Script Date: 12/7/2018 2:16:20 PM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE PROCEDURE [dbo].[Sp_GetAllEmployee]
AS
BEGIN
	SELECT id
      ,userid
      ,firstname
      ,lastname
      ,jobtitle
      ,salary
      ,created_date
      ,created_by
      ,updated_date
      ,updated_by
  FROM tbl_employee
END
GO
