\NeedsTeXFormat{LaTeX2e}
\ProvidesClass{hcmut-report}

\DeclareOption*{\PassOptionsToClass{\CurrentOption}{article}}
\ProcessOptions\relax
\LoadClass[12pt,a4paper]{article}
\usepackage[dvipsnames]{xcolor}

\RequirePackage[utf8]{vntex,inputenc}
\RequirePackage{anyfontsize,amsmath,amssymb}
\RequirePackage[T5]{fontenc}

\RequirePackage{geometry,graphicx,indentfirst,setspace}
\RequirePackage{booktabs}
\RequirePackage[center]{caption}

\RequirePackage{fancyhdr,lastpage,hyperref}
\RequirePackage[super,square,sort]{natbib}
\RequirePackage[nameinlink]{cleveref} % use \Cref{} instead of \ref{}

% Change to '[english,vietnamese]' to use Vietnamese labels
\RequirePackage[english]{babel}

% Custom commands
\newcommand{\Month}{
  \ifcase \the\month
  \or THÁNG 1
  \or THÁNG 2
  \or THÁNG 3
  \or THÁNG 4
  \or THÁNG 5
  \or THÁNG 6
  \or THÁNG 7
  \or THÁNG 8
  \or THÁNG 9
  \or THÁNG 10
  \or THÁNG 11
  \or THÁNG 12
  \fi
}

\newcommand{\@deptname}{Dept name}
\newcommand{\odeptname}[1]{\renewcommand{\@deptname}{#1}}

\newcommand{\@coursename}{Course name}
\newcommand{\ocoursename}[1]{\renewcommand{\@coursename}{#1}}

\newcommand{\@reporttype}{Assignment Report}
\newcommand{\oreporttype}[1]{\renewcommand{\@reporttype}{#1}}

\newcommand{\@students}{Sinh viên thực hiện}
\newcommand{\ostudents}[1]{\renewcommand{\@students}{#1}}

\renewcommand{\@title}{Report title}
\newcommand{\otitle}[1]{\renewcommand{\@title}{#1}}

\newcommand{\@advisor}{Advisor}
\newcommand{\oadvisor}[1]{\renewcommand{\@advisor}{#1}}

\newcommand{\@reportplace}{TP. HỒ CHÍ MINH}
\newcommand{\oreportplace}[1]{\renewcommand{\@reportplace}{#1}}

\newcommand{\@reporttime}{\Month NĂM \the\year}
%\newcommand{\@reporttime}{THÁNG 5 NĂM 2023}
\newcommand{\oreporttime}[1]{\renewcommand{\@reporttime}{#1}}

% Configurations
\geometry{
  paper=a4paper,
  vmargin=2cm,
  hmargin=2cm,
  includeheadfoot=true,
  headheight=31.8pt
}

%\setlength{\headheight}{31.8pt}
%\addtolength{\topmargin}{-8.13pt}

\onehalfspacing%

% Allow page breaks inside align* environment
%\allowdisplaybreaks{}

% Makes a lot of things blue, avoid at all costs
%\everymath{\color{blue}}

% Table of contents section text
\AtBeginDocument{\renewcommand*\contentsname{Mục lục}}

% References section text
\AtBeginDocument{\renewcommand*\refname{Tài liệu tham khảo}}
% \AtBeginDocument{\renewcommand*\bibname{Tài liệu tham khảo}}

% Numbering format
% {equation}{section} means <section number>.<equation number>
% {equation}{subsection} means <section number>.<subsection number>.<equation number>
%\numberwithin{equation}{section}

% Set depth of numbering for sections and table of contents
%\setcounter{secnumdepth}{3}
%\setcounter{tocdepth}{3}

\hypersetup{
  colorlinks=true,
  linkcolor=black,
  citecolor=red,
  urlcolor=blue
}

\pagestyle{fancy}

\renewcommand{\headrulewidth}{0.3pt}
\fancyhead{}
\fancyhead[L]{
  \begin{tabular}{rl}
    \begin{picture}(25,15)(0,0)
      \put(0,-8){\includegraphics[width=8mm,height=8mm]{graphics/hcmut.png}}
    \end{picture}
    &
    \begin{tabular}{l}
      \textbf{\bf \ttfamily Trường Đại học Bách Khoa TP.HCM} \\
      \textbf{\bf \ttfamily Khoa Khoa học và Kỹ thuật Máy tính} \\
    \end{tabular}
  \end{tabular}
}

\renewcommand{\footrulewidth}{0.3pt}
\fancyfoot{}
\fancyfoot[L]{\scriptsize \ttfamily \@reporttype}
\fancyfoot[R]{\scriptsize \ttfamily Page {\thepage}/\pageref{LastPage}}

% Custom pages and sections
\def\coverpage{
  \thispagestyle{empty}

  \vspace*{-3\baselineskip}
  \begin{center}
    ĐẠI HỌC QUỐC GIA THÀNH PHỐ HỒ CHÍ MINH

    TRƯỜNG ĐẠI HỌC BÁCH KHOA

    \MakeUppercase{\@deptname}
    \\[2\baselineskip]
    \includegraphics[scale=0.4]{graphics/hcmut.png}
    \\[2\baselineskip]
    \begin{tabular}{c}
      \textbf{\Large \@coursename}                                 \\
      {}                                                           \\
      \midrule                                                     \\
     % {}                                                           \\
      \textbf{\Large \@reporttype}                                 \\
      {}                                                           \\
      \textbf{\Huge \parbox[c]{0.8\linewidth}{\centering \@title}} \\
      {}                                                           \\
      \bottomrule
    \end{tabular}
    \vfill
    \begin{tabular}{rll}
      {\large Giảng viên hướng dẫn:} & {\large \@advisor} & \\
      {\large Sinh viên thực hiện:}  & {\large Nhóm}      & \\
      %Student: & Name      & ID \\
    \end{tabular}
    \\[3\baselineskip]
    \MakeUppercase{\@reportplace, \@reporttime}
  \end{center}

  \clearpage
}
