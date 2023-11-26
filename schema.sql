create table if not exists User (
   UserID              integer  primary key autoincrement,
   Email               text not null,
   PasswordSaltAndHash text not null
);

create table if not exists SubscriptionType (
    SubscriptionTypeID integer  primary key autoincrement,
    Type               text not null,
    Fee                int  not null
);

create table if not exists Subscription (
    SubscriptionID     integer  primary key autoincrement,
    UserID             int  not null,
    SubscriptionTypeID int  not null,
    Active             int  not null,
    LastInvoiceDate    date not null,
    SignUpDate         date not null,
    
    foreign key (UserID) references User (UserID),
    foreign key (SubscriptionTypeID) references SubscriptionType (SubscriptionTypeID)
);

create table if not exists Word (
    WordID  integer  primary key autoincrement,
    Letters text not null,
    WordOfDayDate date
);

create table if not exists Game (
    GameID        integer  primary key autoincrement,
    WordID        int  not null,
    WinOrLoss     int  not null,
    
    foreign key (WordID) references Word (WordID)
);

create table if not exists UserGame (
    UserGameID integer primary key autoincrement,
    GameID int not null,
    UserID int not null,
    Date date not null,

    foreign key (GameID) references Game (GameID),
    foreign key (UserID) references User (UserID)
);

create table if not exists UserMove (
    UserMoveID integer  primary key autoincrement,
    UserID     int  not null,
    GameID     int  not null,
    Guess      text not null,
    
    foreign key (UserID) references User (UserID),
    foreign key (GameID) references Game (GameID)
);