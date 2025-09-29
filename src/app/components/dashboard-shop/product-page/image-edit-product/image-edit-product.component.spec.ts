import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ImageEditProductComponent } from './image-edit-product.component';

describe('ImageEditProductComponent', () => {
  let component: ImageEditProductComponent;
  let fixture: ComponentFixture<ImageEditProductComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [ImageEditProductComponent]
    });
    fixture = TestBed.createComponent(ImageEditProductComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
