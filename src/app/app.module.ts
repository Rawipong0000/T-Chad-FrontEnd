import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import { HTTP_INTERCEPTORS } from '@angular/common/http';
import { AuthInterceptor } from './auth.interceptor';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { RegisterPopUpComponent } from './components/register-pop-up/register-pop-up.component';
import { HeaderComponent } from './components/header/header.component';
import { DashboardShopComponent } from './components/dashboard-shop/dashboard-shop.component';
import { UserLoginComponent } from './components/user-login/user-login.component';
import { SidebarSearchComponent } from './components/dashboard-shop/sidebar-search/sidebar-search.component';
import { ProductPageComponent } from './components/dashboard-shop/product-page/product-page.component';
import { PurchasePageComponent } from './components/purchase-page/purchase-page.component';
import { MyshopPageComponent } from './components/myshop-page/myshop-page.component';
import { SidebarMyshopComponent } from './components/myshop-page/sidebar-myshop/sidebar-myshop.component';
import { ShopnameEditComponent } from './components/myshop-page/sidebar-myshop/shopname-edit/shopname-edit.component';
import { AddProductPageComponent } from './components/myshop-page/add-product-page/add-product-page.component';
import { EditProductComponent } from './components/myshop-page/edit-product/edit-product.component';
import { EditProductImageComponent } from './components/myshop-page/edit-product/edit-product-image/edit-product-image.component';
import { EditShopnameComponent } from './components/dashboard-shop/sidebar-search/edit-shopname/edit-shopname.component';
import { ImageEditProductComponent } from './components/dashboard-shop/product-page/image-edit-product/image-edit-product.component';
import { OrdersManagePageComponent } from './components/myshop-page/orders-manage-page/orders-manage-page.component';
import { ProfilePageComponent } from './components/profile-page/profile-page.component';
import { EditProfileComponent } from './components/profile-page/edit-profile/edit-profile.component';
import { HistoryPageComponent } from './components/history-page/history-page.component';
import { EditAddressComponent } from './components/profile-page/edit-address/edit-address.component';
import { PromoCodePageComponent } from './components/myshop-page/promo-code-page/promo-code-page.component';
import { AddPromoCodePageComponent } from './components/myshop-page/promo-code-page/add-promo-code-page/add-promo-code-page.component';

@NgModule({
  declarations: [
    AppComponent,
    RegisterPopUpComponent,
    HeaderComponent,
    DashboardShopComponent,
    UserLoginComponent,
    SidebarSearchComponent,
    ProductPageComponent,
    PurchasePageComponent,
    MyshopPageComponent,
    SidebarMyshopComponent,
    ShopnameEditComponent,
    AddProductPageComponent,
    EditProductComponent,
    EditProductImageComponent,
    EditShopnameComponent,
    ImageEditProductComponent,
    OrdersManagePageComponent,
    ProfilePageComponent,
    EditProfileComponent,
    HistoryPageComponent,
    EditAddressComponent,
    PromoCodePageComponent,
    AddPromoCodePageComponent
  ],
  imports: [
    BrowserModule,
    FontAwesomeModule,
    AppRoutingModule,
    FormsModule,
    HttpClientModule
  ],
  providers: [ 
    {
      provide: HTTP_INTERCEPTORS,
      useClass: AuthInterceptor,
      multi: true}
    ],
  bootstrap: [AppComponent]
})
export class AppModule { }
