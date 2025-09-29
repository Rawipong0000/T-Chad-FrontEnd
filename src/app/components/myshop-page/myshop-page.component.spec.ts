import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MyshopPageComponent } from './myshop-page.component';

describe('MyshopPageComponent', () => {
  let component: MyshopPageComponent;
  let fixture: ComponentFixture<MyshopPageComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [MyshopPageComponent]
    });
    fixture = TestBed.createComponent(MyshopPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
