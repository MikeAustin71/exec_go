<%@ Page Language="C#" AutoEventWireup="true" CodeFile="Default.aspx.cs" Inherits="_Default" %>

<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.1//EN" "http://www.w3.org/TR/xhtml11/DTD/xhtml11.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head runat="server">
    <title>Paladin - ACS</title>
    <link href="locStyleSheets/SS_MainSplash.css" rel="stylesheet" type="text/css" />
</head>
<body>
    <form id="MainTitle" runat="server">
        <div>
            <asp:Label ID="labelTitlePaladin_01" runat="server" CSSClass="TitlePaladin_01" Text="Paladin Advanced Computing Systems"></asp:Label>
        </div>
        <div>
             <asp:Label ID="labelEst1994" runat="server" CssClass="Est1994" Text="Est. 1994"></asp:Label>
        </div>
        <div>
             <asp:Label ID="labelTitlePaladin_02" runat="server" CssClass="TitlePaladin_02" Text="Software Engineering For The 21st Century"></asp:Label>
        </div>
        <div id="divPalPic" class="clsPalPic">
             <asp:ImageButton ID="btnPaladinPics"  ToolTip="Enter - Click To Continue" runat="server" ImageUrl="~/Images/PaladinIntroGraphic800.gif"
                OnClick="btnPaladinPics_Click" />
        </div>
        
        <div class="ContanBottomSec">
        
           <div class="ContainBotText">
                
                <div id="divBldTomorrow" class="BldTom">Building A Better Tomorrow...</div>
                
                <div id="divOneLine" class="OneLineCode">One Line Of Code At A Time.</div>
           
           </div>
           
           <div class="ContainKightPic">
           
            <div id="divKnight" class="clsKnightPic">
                 <asp:ImageButton ID="imgKnight" ToolTip="Enter - Click To Continue" runat="server" ImageUrl="~/Images/HorseAlone2.gif"
                    OnClick="btnKnightPic_Click" />
            </div>

          </div>
          
        </div>
        
        <div>
        <asp:ScriptManager ID="ScriptManager1" runat="server" />
        </div>
        </form>
</body>
</html>
