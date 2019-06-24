package models

import (
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    "time"
    "fmt"
    _ "github.com/lib/pq"
)

var (
    DB_PREFIX string
    DB_USER string
    DB_PASSWORD string
    DB_HOST string
    DB_NAME string
    DB_PORT string
    Runmode string
)

type Currency struct {
    Id              int64
    Code            string
    Name            string
    Description     string
    Type            string
    CountryId       int64
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type Order struct {
    Id              int64
    SellerUserId    int64
    BuyerUserId     int64
    CurrencyId      int64
    RateCurrencyId  int64
    Amount          float64
    Rate            float64
    TotalAmount     float64
    SellerCurrencyWalletId      int64
    SellerRateCurrencyWalletId  int64
    BuyerCurrencyWalletId       int64
    BuyerRateCurrencyWalletId   int64
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type OrderBuy struct {
    Id                      int64
    UserId                  int64
    CurrencyId              int64
    RateCurrencyId          int64
    Amount                  float64
    Rate                    float64
    TotalAmount             float64
    CurrencyWalletId        int64
    RateCurrencyWalletId    int64
    Lock                    bool
    CreatedAt               time.Time   `orm:"auto_now_add;type(datetime)"`
}

type OrderCurrency struct {
    Id              int64
    UserId          int64
    CurrencyId      int64
    RateCurrencyId  int64
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type OrderGraph1m struct {
    Id                    int64
    LastOrderId           int64
    CurrencyId            int64
    RateCurrencyId        int64
    Open                  float64
    High                  float64
    Low                   float64
    Close                 float64
    Volume                float64
    Split                 float64
    Dividend              float64
    AbsoluteChange        float64
    PercentChange         float64
    Date                  time.Time   `orm:"auto_now_add;type(datetime)"`
}

type OrderGraph5m struct {
    Id                    int64
    LastOrderId           int64
    CurrencyId            int64
    RateCurrencyId        int64
    Open                  float64
    High                  float64
    Low                   float64
    Close                 float64
    Volume                float64
    Split                 float64
    Dividend              float64
    AbsoluteChange        float64
    PercentChange         float64
    Date                  time.Time   `orm:"auto_now_add;type(datetime)"`
}

type OrderGraph15m struct {
    Id                    int64
    LastOrderId           int64
    CurrencyId            int64
    RateCurrencyId        int64
    Open                  float64
    High                  float64
    Low                   float64
    Close                 float64
    Volume                float64
    Split                 float64
    Dividend              float64
    AbsoluteChange        float64
    PercentChange         float64
    Date                  time.Time   `orm:"auto_now_add;type(datetime)"`
}

type OrderGraph30m struct {
    Id                    int64
    LastOrderId           int64
    CurrencyId            int64
    RateCurrencyId        int64
    Open                  float64
    High                  float64
    Low                   float64
    Close                 float64
    Volume                float64
    Split                 float64
    Dividend              float64
    AbsoluteChange        float64
    PercentChange         float64
    Date                  time.Time   `orm:"auto_now_add;type(datetime)"`
}

type OrderGraph1h struct {
    Id                    int64
    LastOrderId           int64
    CurrencyId            int64
    RateCurrencyId        int64
    Open                  float64
    High                  float64
    Low                   float64
    Close                 float64
    Volume                float64
    Split                 float64
    Dividend              float64
    AbsoluteChange        float64
    PercentChange         float64
    Date                  time.Time   `orm:"auto_now_add;type(datetime)"`
}

type OrderGraph6h struct {
    Id                    int64
    LastOrderId           int64
    CurrencyId            int64
    RateCurrencyId        int64
    Open                  float64
    High                  float64
    Low                   float64
    Close                 float64
    Volume                float64
    Split                 float64
    Dividend              float64
    AbsoluteChange        float64
    PercentChange         float64
    Date                  time.Time   `orm:"auto_now_add;type(datetime)"`
}

type OrderGraph12h struct {
    Id                    int64
    LastOrderId           int64
    CurrencyId            int64
    RateCurrencyId        int64
    Open                  float64
    High                  float64
    Low                   float64
    Close                 float64
    Volume                float64
    Split                 float64
    Dividend              float64
    AbsoluteChange        float64
    PercentChange         float64
    Date                  time.Time   `orm:"auto_now_add;type(datetime)"`
}

type OrderGraph1d struct {
    Id                    int64
    LastOrderId           int64
    CurrencyId            int64
    RateCurrencyId        int64
    Open                  float64
    High                  float64
    Low                   float64
    Close                 float64
    Volume                float64
    Split                 float64
    Dividend              float64
    AbsoluteChange        float64
    PercentChange         float64
    Date                  time.Time   `orm:"auto_now_add;type(datetime)"`
}

type OrderGraph7d struct {
    Id                    int64
    LastOrderId           int64
    CurrencyId            int64
    RateCurrencyId        int64
    Open                  float64
    High                  float64
    Low                   float64
    Close                 float64
    Volume                float64
    Split                 float64
    Dividend              float64
    AbsoluteChange        float64
    PercentChange         float64
    Date                  time.Time   `orm:"auto_now_add;type(datetime)"`
}

type OrderSell struct {
    Id                      int64
    UserId                  int64
    CurrencyId              int64
    RateCurrencyId          int64
    Amount                  float64
    Rate                    float64
    TotalAmount             float64
    CurrencyWalletId        int64
    RateCurrencyWalletId    int64
    Lock                    bool
    CreatedAt               time.Time   `orm:"auto_now_add;type(datetime)"`
}

type User struct {
    Id              int64
    Name            string      `valid:"Required;MaxSize(250)"`
    Username        string      `valid:"Required;MaxSize(250)"`
    EditNameTimes   int
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type Wallet struct {
    Id              int64
    WalletMasterId  int64
    Amount          float64
    AmountLocked    float64
    Nickname        string
    Primary         bool
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type WalletCrypto struct {
    Id              int64
    WalletId        int64
    Address         string
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type WalletMaster struct {
    Id              int64
    UserId          int64
    CurrencyId      int64
    CurrencyType    string
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type WalletPassphrase struct {
    Id              int64
    WalletId        int64
    Passphrase      string
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

func init() {
    DB_PREFIX = beego.AppConfig.String("DB_PREFIX")
    DB_USER = beego.AppConfig.String("DB_USER")
    DB_PASSWORD = beego.AppConfig.String("DB_PASSWORD")
    DB_HOST = beego.AppConfig.String("DB_HOST")
    DB_NAME = beego.AppConfig.String("DB_NAME")
    DB_PORT = beego.AppConfig.String("DB_PORT")
    Runmode = beego.AppConfig.String("runmode")

    switch Runmode {
    case "dev":
        orm.RegisterDriver("postgres", orm.DRPostgres)
        orm.RegisterDataBase("default", "postgres", fmt.Sprintf("postgres://%s:%s@%s/%s?port=%i&sslmode=disable",
        DB_USER, DB_PASSWORD, DB_HOST, DB_NAME, DB_PORT), 30)
        //orm.Debug = true
    }
}